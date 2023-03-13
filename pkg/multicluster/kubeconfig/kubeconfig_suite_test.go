package kubeconfig_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKubeconfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kubeconfig Suite")
}
