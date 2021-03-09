package e2e

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestE2e(t *testing.T) {
	initVars()
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}
