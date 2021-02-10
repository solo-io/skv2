package main

import (
	"fmt"
	"os"

	"github.com/solo-io/go-list-licenses/pkg/license"
)

func main() {
	packages := []string{
		"github.com/solo-io/skv2/pkg/resource",
		"github.com/solo-io/skv2/pkg/utils",
		"github.com/solo-io/skv2/pkg/source",
		"github.com/solo-io/skv2/pkg/stats",
		"github.com/solo-io/skv2/pkg/verifier",
		"github.com/solo-io/skv2/pkg/workqueue",
	}

	// dependencies for this package which are used on mac, and will not be present in linux CI
	macOnlyDependencies := []string{
		"github.com/mitchellh/go-homedir",
		"github.com/containerd/continuity",
		"golang.org/x/sys/unix",
	}

	app := license.Cli(packages, macOnlyDependencies)
	if err := app.Execute(); err != nil {
		fmt.Errorf("unable to run oss compliance check: %v\n", err)
		os.Exit(1)
	}
}
