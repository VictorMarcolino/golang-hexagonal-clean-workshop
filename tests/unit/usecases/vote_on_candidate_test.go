package usecases_test

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/adapters"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/tests/specs"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("USECASE:IN-MEMORY:VoteOnCandidateUsecase", func() {
	adapter := adapters.NewInMemoryAdapter()
	voteOnCandidateUsecase := usecases.NewVoteOnCandidateUsecase(adapter)

	specs.VoteOnCandidateUsecase(voteOnCandidateUsecase, adapter)
})
