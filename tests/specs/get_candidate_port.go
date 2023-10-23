package specs

import (
	"errors"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domainerrors"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type getCandidatePortTestRequires interface {
	ports.GetCandidateI
	ports.CreateCandidateI
}

func GetCandidatePort(implementation getCandidatePortTestRequires) {
	When("Given port is used", func() {
		var validUUID string
		expectedCandidate := domain.Candidate{Name: "Valid Candidate"}

		BeforeEach(func() {
			var err error
			validUUID, err = implementation.CreateCandidate(expectedCandidate)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("When a valid UUID is provided", func() {
			It("Should be able to fetch the corresponding candidate", func() {
				candidate, err := implementation.GetCandidate(validUUID)
				Expect(err).ToNot(HaveOccurred())
				expectedCandidate.UUID = validUUID
				Expect(candidate).To(Equal(expectedCandidate))
			})
		})

		Context("When a nonexistent UUID is provided", func() {
			nonexistentUUID := "12345-nonexistent"

			It("Should raise an error", func() {
				_, err := implementation.GetCandidate(nonexistentUUID)
				Expect(err).To(HaveOccurred())
				Expect(errors.Is(err, domainerrors.ErrCandidateNotFound)).To(BeTrue())
			})
		})
	})
}
