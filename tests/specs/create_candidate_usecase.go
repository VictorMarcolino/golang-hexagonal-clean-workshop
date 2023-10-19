package specs

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type CreateCandidateUsecaseTestRequires interface {
	usecases.CreateCandidateUsecaseI
}

func CreateCandidateUsecase(useCases CreateCandidateUsecaseTestRequires) {
	When("Given name was provided", func() {
		testResource := domain.Candidate{Name: "Provided"}
		It("Expect usecase to not raise an error", func() {
			_, err := useCases.ExecuteCreateCandidate(testResource)
			Expect(err).ToNot(HaveOccurred())
		})
		It("Expect usecase to return an id", func() {
			uuid, _ := useCases.ExecuteCreateCandidate(testResource)
			Expect(uuid).ToNot(BeEmpty())
		})
	})
	When("Name was not provided", func() {
		testResource := domain.Candidate{Name: ""}
		It("Expect usecase to raise a error", func() {
			_, err := useCases.ExecuteCreateCandidate(testResource)
			Expect(err).To(HaveOccurred())
		})
		It("Expect usecase to not return an id", func() {
			uuid, _ := useCases.ExecuteCreateCandidate(testResource)
			Expect(uuid).To(BeEmpty())
		})
	})
}
