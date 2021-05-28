package model

import (
	"fmt"
	"strings"

	"github.com/solo-io/skv2/codegen/util/stringutils"
)

var (
	// function for determining the relative path of generated api types package
	TypesRelativePath = func(kind, version string) string {
		return fmt.Sprintf("pkg/apis/%v/%v", strings.ToLower(stringutils.Pluralize(kind)), version)
	}

	// function for determining the relative path of generated schduler package
	SchedulerRelativePath = "pkg/scheduler"

	// function for determining the relative path of generated finalizer package
	FinalizerRelativePath = "pkg/finalizer"

	// function for determining the relative path of generated parameters package
	ParametersRelativePath = "pkg/parameters"

	// function for determining the relative path of generated metrics package
	MetricsRelativePath = "pkg/metrics"
)
