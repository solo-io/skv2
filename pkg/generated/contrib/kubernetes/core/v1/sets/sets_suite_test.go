package v1_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSets(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sets Suite")
}
