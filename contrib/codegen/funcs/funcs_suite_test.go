package funcs_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFuncs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Funcs Suite")
}
