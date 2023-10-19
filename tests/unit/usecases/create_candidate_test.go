package usecases_test

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

var _ = Describe("CreateCandidateUsecase", func() {
	specs.CreateCandidateUsecase(nil)
})
