package testmongodbhandler_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestmongodbhandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testmongodbhandler Suite")
}
