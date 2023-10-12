package unit_test

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/tests/specs"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSuit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestSuit:TestSuit")
}

var _ = Describe("Describe", func() {
	specs.ExampleSpec()
})
