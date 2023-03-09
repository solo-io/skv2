package verifier_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestVerifier(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Verifier Suite")
}
