package codegen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/openapi"
	"cuelang.org/go/encoding/protobuf"
	"github.com/sirupsen/logrus"
	"github.com/solo-io/anyvendor/anyvendor"
	"github.com/solo-io/anyvendor/pkg/manager"
	"github.com/solo-io/skv2/builder"
	"github.com/solo-io/skv2/codegen/collector"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/codegen/proto"
	"github.com/solo-io/skv2/codegen/render"
	"github.com/solo-io/skv2/codegen/skv2_anyvendor"
	"github.com/solo-io/skv2/codegen/util"
	"github.com/solo-io/skv2/codegen/writer"
)

const (
	DefaultHeader = "Code generated by skv2. DO NOT EDIT."
)

// runs the codegen compilation for the current Go module
type Command struct {
	// the name of the app or component
	// used to label k8s manifests
	AppName string

	// config to vendor protos and other non-go files
	// Optional: If nil will not be used
	AnyVendorConfig *skv2_anyvendor.Imports

	// the k8s api groups for which to compile
	Groups []render.Group

	// optinal helm chart to render
	Chart *model.Chart

	// the root directory for generated Kube manifests
	ManifestRoot string

	// optional Go/Docker images to build
	Builds []model.Build

	// the root directory for Build files (Dockerfile, entrypoint script, etc.)
	BuildRoot string

	// Header to prepend to generated files.
	// If not set will default to: "Code generated by skv2. DO NOT EDIT."
	// The string will auto prepend the comment character depending on the file type it is being written to.
	GeneratedHeader string

	// custom builder to shim Go Build and Docker Build commands (for testing)
	// If not provided, skv2 will exec
	// go and docker commands
	Builder builder.Builder

	// the path to the root dir of the module on disk
	// files will be written relative to this dir,
	// except kube clientsets which
	// will generate to the module of the group
	moduleRoot string

	// the name of the go module (as a go package)
	moduleName string

	// context of the command
	ctx context.Context
}

// function to execute skv2 code gen from another repository
func (c Command) Execute() error {
	c.ctx = context.Background()
	c.moduleRoot = util.GetModuleRoot()
	c.moduleName = util.GetGoModule()
	if c.Builder == nil {
		c.Builder = builder.NewBuilder()
	}

	if c.GeneratedHeader == "" {
		c.GeneratedHeader = DefaultHeader
	}

	if err := c.generateChart(); err != nil {
		return err
	}

	for _, group := range c.Groups {
		// init connects children to their parents
		group.Init()

		if err := c.generateGroup(group); err != nil {
			return err
		}
	}

	for _, build := range c.Builds {
		if err := c.generateBuild(build); err != nil {
			return err
		}
		if err := c.buildPushImage(build); err != nil {
			return err
		}
	}
	return nil
}

func (c Command) generateChart() error {
	if c.Chart != nil {
		files, err := render.RenderChart(*c.Chart)
		if err != nil {
			return err
		}

		writer := &writer.DefaultFileWriter{
			Root:   filepath.Join(c.moduleRoot, c.ManifestRoot),
			Header: c.GeneratedHeader,
		}

		if err := writer.WriteFiles(files); err != nil {
			return err
		}
	}

	return nil
}

func (c Command) generateGroup(grp model.Group) error {
	if err := c.compileProtosAndUpdateGroup(&grp); err != nil {
		return err
	}

	fileWriter := &writer.DefaultFileWriter{
		Root:   c.moduleRoot,
		Header: c.GeneratedHeader,
	}

	protoTypes, err := render.RenderProtoTypes(grp)
	if err != nil {
		return err
	}

	if err := fileWriter.WriteFiles(protoTypes); err != nil {
		return err
	}

	apiTypes, err := render.RenderApiTypes(grp)
	if err != nil {
		return err
	}

	if err := fileWriter.WriteFiles(apiTypes); err != nil {
		return err
	}

	manifests, err := render.RenderManifests(c.AppName, c.ManifestRoot, grp)
	if err != nil {
		return err
	}

	if err := fileWriter.WriteFiles(manifests); err != nil {
		return err
	}

	if err := render.KubeCodegen(grp); err != nil {
		return err
	}

	return nil
}

// compiles protos and attaches descriptors to the group and its resources
// it is important to run this func before rendering as it attaches protos to the
// group model
func (c Command) compileProtosAndUpdateGroup(grp *render.Group) error {
	if !grp.RenderProtos {
		return nil
	}

	if grp.ProtoDir == "" {
		grp.ProtoDir = anyvendor.DefaultDepDir
	}

	if c.AnyVendorConfig != nil {
		mgr, err := manager.NewManager(c.ctx, c.moduleRoot)
		if err != nil {
			return err
		}

		if err := mgr.Ensure(c.ctx, c.AnyVendorConfig.ToAnyvendorConfig()); err != nil {
			return err
		}
	}
	descriptors, err := proto.CompileProtos(
		grp.Module,
		grp.ApiRoot,
		grp.ProtoDir,
	)
	if err != nil {
		return err
	}

	// set the descriptors on the group for compilation
	grp.Descriptors = descriptors

	for i, resource := range grp.Resources {
		// attach the proto messages for spec and status to each resource
		// these are processed by renderers at later stages
		addMessagesToResource(descriptors, &resource)
		grp.Resources[i] = resource
	}

	return nil
}

func addMessagesToResource(descriptors []*collector.DescriptorWithPath, resource *model.Resource) {
	var foundSpec bool
	for _, fileDescriptor := range descriptors {

		if fileDescriptor.GetPackage() == resource.Group.Group {

			protoDir := resource.Group.ProtoDir
			if protoDir == "" {
				protoDir = anyvendor.DefaultDepDir
			}
			coll := collector.NewCollector([]string{protoDir}, nil)
			imports, err := coll.CollectImportsForFile(protoDir, fileDescriptor.ProtoFilePath)
			if err != nil {
				log.Fatal(err)
			}
			cfg := &protobuf.Config{
				Root:   resource.Group.ProtoDir,
				Module: resource.Group.Group,
				Paths:  imports,
			}
			ext := protobuf.NewExtractor(cfg)
			if err := ext.AddFile(fileDescriptor.ProtoFilePath, nil); err != nil {
				log.Fatal(err)
			}
			instances, err := ext.Instances()
			if err != nil {
				log.Fatal(err)
			}
			generator := &openapi.Generator{
				ExpandReferences: true,
			}
			built := cue.Build(instances)
			for _, builtInstance := range built {
				if builtInstance.Err != nil {
					log.Fatal(err)
				}
				if err := builtInstance.Value().Validate(); err != nil {
					log.Fatal(err)
				}
				oapi, err := generator.Schemas(builtInstance)
				if err != nil {
					log.Fatal(err)
				}
				byt, err := json.Marshal(oapi)
				if err != nil {
					log.Fatal(err)
				}
				buf := &bytes.Buffer{}
				if err := json.Indent(buf, byt, "", " "); err != nil {
					log.Fatal(err)
				}
				fmt.Println(buf.String())
			}

			if specMessage := fileDescriptor.GetMessage(resource.Spec.Type.Name); specMessage != nil {
				resource.Spec.Type.Message = specMessage
				resource.Spec.Type.GoPackage = fileDescriptor.GetOptions().GetGoPackage()
				foundSpec = true
			}

			if resource.Status != nil {
				if statusMessage := fileDescriptor.GetMessage(resource.Status.Type.Name); statusMessage != nil {
					resource.Status.Type.Message = statusMessage
				}
			}

		}
	}
	if !foundSpec {
		logrus.Warnf("no package found for %v", resource.Group.Group)
	}
}

func (c Command) generateBuild(build model.Build) error {
	buildFiles, err := render.RenderBuild(build)
	if err != nil {
		return err
	}

	writer := &writer.DefaultFileWriter{Root: c.BuildRoot}

	if err := writer.WriteFiles(buildFiles); err != nil {
		return err
	}

	return nil
}

func (c Command) buildPushImage(build model.Build) error {
	ldFlags := fmt.Sprintf("-X %v/pkg/version.Version=%v", c.moduleRoot, build.Tag)

	// get the main package from the main directory
	// assumes package == module name + main dir path
	mainkPkg := filepath.Join(c.moduleName, filepath.Dir(build.MainFile))

	buildDir := filepath.Join(c.BuildRoot, build.Repository)

	binName := filepath.Join(buildDir, build.Repository+"-linux-amd64")

	log.Printf("Building main package at %v ...", mainkPkg)

	err := c.Builder.GoBuild(util.GoCmdOptions{
		BinName: binName,
		Args: []string{
			"-ldflags=" + ldFlags,
			`-gcflags='all="-N -l"''`,
		},
		PackagePath: mainkPkg,
		Env: []string{
			"GO111MODULE=on",
			"CGO_ENABLED=0",
			"GOARCH=amd64",
			"GOOS=linux",
		},
	})
	if err != nil {
		return err
	}

	defer os.Remove(binName)

	fullImageName := fmt.Sprintf("%v/%v:%v", build.Registry, build.Repository, build.Tag)

	log.Printf("Building docker image %v ...", fullImageName)
	if err := c.Builder.Docker("build", "-t", fullImageName, buildDir); err != nil {
		return err
	}

	if !build.Push {
		return nil
	}

	log.Printf("Pushing docker image %v ...", fullImageName)

	return c.Builder.Docker("push", fullImageName)
}
