package crdutils_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCrdutils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Crdutils Suite")
}
