package specs

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type VoteOnCandidateUsecaseTestRequiresUseCases interface {
	usecases.VoteOnCandidateUsecaseI
}

type VoteOnCandidateUsecaseTestRequiresPorts interface {
	ports.GetCandidateI
	ports.GetVotingSessionI
	ports.GetVotesI
	ports.CreateCandidateI
	ports.CreateVotingI
}

func VoteOnCandidateUsecase(useCases VoteOnCandidateUsecaseTestRequiresUseCases, adapters VoteOnCandidateUsecaseTestRequiresPorts) {
	When("Given voting exists", func() {
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
		})

		When("Given Candidate is part of the voting", func() {
			It("Expect usecase to not raise an error", func() {
				err := useCases.ExecuteVoteOnCandidateUsecase(candidate1, votingSession)
				Expect(err).ToNot(HaveOccurred())
			})

			It("Expect vote counting to have increased", func() {
				beforeVoteCount, err := adapters.GetVotes(votingSession.UUID)
				Expect(err).ToNot(HaveOccurred())

				err = useCases.ExecuteVoteOnCandidateUsecase(candidate1, votingSession)
				Expect(err).ToNot(HaveOccurred())

				afterVoteCount, err := adapters.GetVotes(votingSession.UUID)

				Expect(err).ToNot(HaveOccurred())
				Expect(len(afterVoteCount)).To(BeNumerically(">", len(beforeVoteCount)))
			})
		})

		When("Given Candidate is NOT part of the voting", func() {
			It("Expect usecase to raise an error", func() {
				invalidCandidate := domain.Candidate{UUID: "invalidUUID", Name: "Invalid Candidate"}
				err := useCases.ExecuteVoteOnCandidateUsecase(invalidCandidate, votingSession)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	When("Given Voting does NOT exists", func() {
		It("Expect usecase to raise an error", func() {
			nonExistentVoting := domain.Voting{UUID: "nonExistentUUID", Name: "Non Existent Voting"}
			validCandidate := domain.Candidate{UUID: "validUUID", Name: "Valid Candidate"}
			err := useCases.ExecuteVoteOnCandidateUsecase(validCandidate, nonExistentVoting)
			Expect(err).To(HaveOccurred())
		})
	})
}
