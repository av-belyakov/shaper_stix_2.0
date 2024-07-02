package testhandlercasethehive_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTesthandlercasethehive(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testhandlercasethehive Suite")
}
