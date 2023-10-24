package specs

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type CountVotesUsecaseTestRequiresUseCases interface {
	usecases.GetVotingResultUsecaseI
}

type CountVotesUsecaseTestRequiresPorts interface {
	ports.GetVotingSessionI
	ports.GetCandidateI
	ports.CreateVoteI
	ports.CreateVotingI
	ports.CreateCandidateI
}

func CountVotesUsecase(useCases CountVotesUsecaseTestRequiresUseCases, adapters CountVotesUsecaseTestRequiresPorts) {
	When("Given a valid voting session", func() {
		var (
			votingSession domain.Voting
			candidate1    domain.Candidate
			candidate2    domain.Candidate
		)

		BeforeEach(func() {
			var err error
			candidate1 = domain.Candidate{UUID: utils.GenerateUUID(), Name: "Joao"}
			candidate2 = domain.Candidate{UUID: utils.GenerateUUID(), Name: "Maria"}

			uuid1, err := adapters.CreateCandidate(candidate1)
			Expect(err).ToNot(HaveOccurred())
			candidate1.UUID = uuid1

			uuid2, err := adapters.CreateCandidate(candidate2)
			Expect(err).ToNot(HaveOccurred())
			candidate2.UUID = uuid2

			votingSession = domain.Voting{Name: "City Council", Candidates: []domain.Candidate{candidate1, candidate2}}
			votingSessionUuid, err := adapters.CreateVoting(votingSession)
			Expect(err).ToNot(HaveOccurred())
			votingSession.UUID = votingSessionUuid

			_, _ = adapters.CreateVote(domain.Vote{
				VotingUUID:    votingSession.UUID,
				CandidateUuid: candidate2.UUID,
			})
		})

		It("Expect usecase to not raise an error", func() {
			_, err := useCases.ExecuteGetVotingResultUsecase(votingSession.UUID)
			Expect(err).ToNot(HaveOccurred())
		})

		It("Expect to get a valid vote count", func() {
			voteCounts, err := useCases.ExecuteGetVotingResultUsecase(votingSession.UUID)
			Expect(err).ToNot(HaveOccurred())
			Expect(voteCounts).To(BeEquivalentTo(map[string]int{
				candidate2.UUID: 1,
			}))
		})
	})

	When("Given an invalid voting session", func() {
		It("Expect usecase to raise an error", func() {
			_, err := useCases.ExecuteGetVotingResultUsecase("invalidUUID")
			Expect(err).To(HaveOccurred())
		})
	})
}
