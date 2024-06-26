package testreportwrap_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestReportwrap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reportwrap Suite")
}
