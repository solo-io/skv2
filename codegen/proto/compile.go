package proto

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/solo-io/skv2/codegen/collector"
	"github.com/solo-io/skv2/codegen/util"
)

// make sure the pkg matches the go_package option in the proto
// TODO: validate this
func CompileProtos(goModule, moduleName, protoDir string) ([]*collector.DescriptorWithPath, error) {
	log.Printf("Compiling protos in %v", protoDir)

	// need to be in module root so protoc runs on the expecte
	if err := os.Chdir(util.GetModuleRoot()); err != nil {
		return nil, err
	}

	protoDir, err := filepath.Abs(protoDir)
	if err != nil {
		return nil, err
	}
	protoOutDir, err := ioutil.TempDir("", "")
	if err != nil {
		return nil, err
	}
	defer os.Remove(protoOutDir)

	coll := collector.NewCollector(
		nil,
		[]string{protoDir}, // import the inputs dir //
	)

	descriptors, err := collector.NewProtoCompiler(
		coll,
		[]string{protoDir}, // import the inputs dir
		nil,
		[]string{
			"jsonshim",
		},
		protoOutDir,
		func(file string) bool {
			return true
		}).CompileDescriptorsFromRoot(protoDir, nil)
	if err != nil {
		return nil, err
	}

	// copy the files generated for our package into our repo from the
	// tmp dir
	return descriptors, copyFiles(filepath.Join(protoOutDir, moduleName), goModule)
}

func copyFiles(srcDir, destDir string) error {
	if err := filepath.Walk(srcDir, func(srcFile string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		destFile := filepath.Join(destDir, strings.TrimPrefix(srcFile, srcDir))

		// copy
		srcReader, err := os.Open(srcFile)
		if err != nil {
			return err
		}
		defer srcReader.Close()

		if err := os.MkdirAll(filepath.Dir(destFile), 0755); err != nil {
			return err
		}

		dstFile, err := os.Create(destFile)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		log.Printf("copying %v -> %v", srcFile, destFile)
		_, err = io.Copy(dstFile, srcReader)
		return err

	}); err != nil {
		return err
	}

	return nil
}
