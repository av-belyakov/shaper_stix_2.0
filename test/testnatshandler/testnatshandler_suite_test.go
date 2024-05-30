package testnatshandler_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestnatshandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testnatshandler Suite")
}
