package usecases_test

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/adapters"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/tests/specs"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("USECASE:IN-MEMORY:CreateVotingUsecase", func() {
	adapter := adapters.NewInMemoryAdapter()
	createVotingUsecase := usecases.NewCreateVotingUsecase(adapter)

	specs.CreateVotingUsecase(createVotingUsecase, adapter)
})
