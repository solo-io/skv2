package contrib_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestContrib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contrib Suite")
}
