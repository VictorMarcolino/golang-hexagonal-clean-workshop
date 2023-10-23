package filesystemadapter_test

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/adapters/filesystemadapter"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/tests/specs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSuit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PORT:FILESYSTEM")
}

var _ = Describe("PORT:FILESYSTEM:CreateCandidate", func() {
	adapter := filesystemadapter.NewFileSystemAdapter()
	specs.CreateCandidatePort(adapter)
	specs.CreateVotingPort(adapter)
	specs.CreateVotePort(adapter)
	specs.GetVotesPort(adapter)
	specs.GetCandidatePort(adapter)
	specs.GetVotingSessionPort(adapter)
})
