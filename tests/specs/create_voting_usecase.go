package specs

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type CreateVotingUsecaseTestRequiresUseCases interface {
	usecases.CreateVotingUsecaseI
}

type CreateVotingUsecaseTestRequiresPorts interface {
	ports.CreateCandidateI
	ports.GetVotingSessionI
}

func CreateVotingUsecase(useCases CreateVotingUsecaseTestRequiresUseCases, adapters CreateVotingUsecaseTestRequiresPorts) {
	When("Given 2 candidates exist", func() {
		var (
			candidates []domain.Candidate
		)
		BeforeEach(func() {
			var err error
			candidate1 := domain.Candidate{UUID: utils.GenerateUUID(), Name: "Joao"}
			candidate2 := domain.Candidate{UUID: utils.GenerateUUID(), Name: "Maria"}

			uuid1, err := adapters.CreateCandidate(candidate1)
			Expect(err).ToNot(HaveOccurred())
			candidate1.UUID = uuid1

			uuid2, err := adapters.CreateCandidate(candidate2)
			Expect(err).ToNot(HaveOccurred())
			candidate2.UUID = uuid2

			candidates = []domain.Candidate{candidate1, candidate2}
		})

		It("Expect usecase to not raise an error", func() {
			_, err := useCases.ExecuteCreateVotingUsecase("City Council", candidates)
			Expect(err).ToNot(HaveOccurred())
		})

		It("Expect a new voting session to be created", func() {
			votingUUID, err := useCases.ExecuteCreateVotingUsecase("City Council", candidates)
			Expect(err).ToNot(HaveOccurred())

			_, err = adapters.GetVotingSession(votingUUID)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	When("Given an empty list of candidates", func() {
		It("Expect usecase to raise an error", func() {
			_, err := useCases.ExecuteCreateVotingUsecase("City Council", []domain.Candidate{})
			Expect(err).To(HaveOccurred())
		})
	})
}
