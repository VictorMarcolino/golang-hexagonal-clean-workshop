package usecases_test

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/tests/specs"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("CreateCandidateUsecase", func() {
	specs.CreateCandidateUsecase(nil)
})
