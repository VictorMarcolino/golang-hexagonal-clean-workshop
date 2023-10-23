package specs

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type createVotePortTestRequires interface {
	ports.CreateVoteI
}

func CreateVotePort(implementation createVotePortTestRequires) {
	When("Given port is used", func() {
		var uuid string
		testResource := domain.Vote{VotingUUID: "validVotingUUID", CandidateUuid: "validCandidateUUID"}

		Context("When a valid vote is provided", func() {
			It("Expect usecase to not raise an error", func() {
				var err error
				uuid, err = implementation.CreateVote(testResource)
				Expect(err).ToNot(HaveOccurred())
				Expect(uuid).ToNot(BeEmpty())
			})
		})
	})
}
