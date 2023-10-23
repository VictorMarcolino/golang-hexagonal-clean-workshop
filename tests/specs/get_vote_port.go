package specs

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type getVotesPortTestRequires interface {
	ports.GetVotesI
	ports.CreateVoteI
}

func GetVotesPort(implementation getVotesPortTestRequires) {
	When("Given port is and there is already data to be returned", func() {
		validVotingUUID := utils.GenerateUUID()
		var expectedVotes []domain.Vote

		BeforeEach(func() {
			var err error
			expectedVotes = []domain.Vote{
				{VotingUUID: validVotingUUID, CandidateUuid: "candidate1"},
				{VotingUUID: validVotingUUID, CandidateUuid: "candidate2"},
			}
			for i := range expectedVotes {
				expectedVotes[i].UUID, err = implementation.CreateVote(expectedVotes[i])
				Expect(err).ToNot(HaveOccurred())
			}
		})

		Context("When a valid voting UUID is provided", func() {
			It("Should return the corresponding votes", func() {
				votes, err := implementation.GetVotes(validVotingUUID)
				Expect(err).ToNot(HaveOccurred())
				Expect(votes).To(ConsistOf(expectedVotes))
			})
		})
	})
}
