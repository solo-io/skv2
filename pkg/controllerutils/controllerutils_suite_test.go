package controllerutils_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestControllerutils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllerutils Suite")
}
