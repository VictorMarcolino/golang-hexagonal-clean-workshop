package ports_test

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/adapters"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/tests/specs"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("PORT:IN-MEMORY:GetVotingSessionPort", func() {
	adapter := adapters.NewInMemoryAdapter()
	specs.GetVotingSessionPort(adapter)
})
