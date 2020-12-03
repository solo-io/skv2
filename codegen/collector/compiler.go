package collector

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/log"
	"golang.org/x/sync/errgroup"
)

func NewProtoCompiler(collector Collector, customImports, customGogoArgs, customPlugins []string,
	descriptorOutDir string, wantCompile func(string) bool) *protoCompiler {
	return &protoCompiler{
		collector:        collector,
		descriptorOutDir: descriptorOutDir,
		customImports:    customImports,
		customArgs:       customGogoArgs,
		wantCompile:      wantCompile,
		customPlugins:    customPlugins,
	}
}

type ProtoCompiler interface {
	CompileDescriptorsFromRoot(root string, skipDirs []string) ([]*DescriptorWithPath, error)
}

type protoCompiler struct {
	collector        Collector
	descriptorOutDir string
	customImports    []string
	customArgs       []string
	wantCompile      func(string) bool
	customPlugins    []string
}

func (p *protoCompiler) CompileDescriptorsFromRoot(root string, skipDirs []string) ([]*DescriptorWithPath, error) {
	var descriptors []*DescriptorWithPath
	var mutex sync.Mutex
	addDescriptor := func(f DescriptorWithPath) {
		mutex.Lock()
		defer mutex.Unlock()
		descriptors = append(descriptors, &f)
	}
	var group errgroup.Group
	for _, dir := range append([]string{root}) {
		absoluteDir, err := filepath.Abs(dir)
		if err != nil {
			return nil, err
		}
		walkErr := filepath.Walk(absoluteDir, func(protoFile string, info os.FileInfo, err error) error {
			if !strings.HasSuffix(protoFile, ".proto") {
				return nil
			}
			for _, skip := range skipDirs {
				skipRoot := filepath.Join(absoluteDir, skip)
				if strings.HasPrefix(protoFile, skipRoot) {
					log.Warnf("skipping proto %v because it is %v is a skipped directory", protoFile, skipRoot)
					return nil
				}
			}

			// parallelize parsing the descriptors as each one requires file i/o and is slow
			group.Go(func() error {
				imports, err := p.collector.CollectImportsForFile(absoluteDir, protoFile)
				if err != nil {
					return err
				}
				return p.addDescriptorsForFile(addDescriptor, imports, protoFile)
			})
			return nil
		})
		if walkErr != nil {
			return nil, walkErr
		}

		// Wait for all descriptor parsing to complete.
		if err := group.Wait(); err != nil {
			return nil, err
		}
	}
	sort.SliceStable(descriptors, func(i, j int) bool {
		return descriptors[i].GetName() < descriptors[j].GetName()
	})

	// don't add the same proto twice, this avoids the issue where a dependency is imported multiple times
	// with different import paths
	return filterDuplicateDescriptors(descriptors), nil
}

func (p *protoCompiler) addDescriptorsForFile(addDescriptor func(f DescriptorWithPath), imports []string, protoFile string) error {
	log.Printf("processing proto file input %v", protoFile)
	// don't generate protos for non-project files
	compile := p.wantCompile(protoFile)

	// use a temp file to store the output from protoc, then parse it right back in
	// this is how we "wrap" protoc
	tmpFile, err := ioutil.TempFile("", "solo-kit-gen-")
	if err != nil {
		return err
	}
	if err := tmpFile.Close(); err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	if err := p.writeDescriptors(protoFile, tmpFile.Name(), imports, compile); err != nil {
		return eris.Wrapf(err, "writing descriptors")
	}
	desc, err := p.readDescriptors(tmpFile.Name())
	if err != nil {
		return eris.Wrapf(err, "reading descriptors")
	}

	for _, f := range desc.File {
		descriptorWithPath := DescriptorWithPath{FileDescriptorProto: f}
		if strings.HasSuffix(protoFile, f.GetName()) {
			descriptorWithPath.ProtoFilePath = protoFile
		}
		addDescriptor(descriptorWithPath)
	}

	return nil
}

var defaultGogoArgs = []string{
	"plugins=grpc",
}

func (p *protoCompiler) writeDescriptors(protoFile, toFile string, imports []string, compileProtos bool) error {
	cmd := exec.Command("protoc")
	for i := range imports {
		imports[i] = "-I" + imports[i]
	}
	cmd.Args = append(cmd.Args, imports...)
	gogoArgs := append(defaultGogoArgs, p.customArgs...)

	if compileProtos {
		cmd.Args = append(cmd.Args,
			"--go_out="+strings.Join(gogoArgs, ",")+":"+p.descriptorOutDir,
			"--ext_out="+strings.Join(gogoArgs, ",")+":"+p.descriptorOutDir,
		)

		for _, plugin := range p.customPlugins {
			cmd.Args = append(cmd.Args,
				"--"+plugin+"_out="+strings.Join(gogoArgs, ",")+":"+p.descriptorOutDir,
			)
		}
	}

	cmd.Args = append(cmd.Args, "-o"+toFile, "--include_imports", "--include_source_info",
		protoFile)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return eris.Wrapf(err, "%v failed: %s", cmd.Args, out)
	}
	return nil
}

func (p *protoCompiler) readDescriptors(fromFile string) (*descriptor.FileDescriptorSet, error) {
	var desc descriptor.FileDescriptorSet
	protoBytes, err := ioutil.ReadFile(fromFile)
	if err != nil {
		return nil, eris.Wrapf(err, "reading file")
	}
	if err := proto.Unmarshal(protoBytes, &desc); err != nil {
		return nil, eris.Wrapf(err, "unmarshalling tmp file as descriptors")
	}
	return &desc, nil
}
