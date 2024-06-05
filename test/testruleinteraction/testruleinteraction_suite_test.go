package testruleinteraction_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestruleinteraction(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testruleinteraction Suite")
}
