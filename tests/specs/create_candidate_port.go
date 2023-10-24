package specs

import (
	"errors"

	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domainerrors"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type createCandidatePortTestRequires interface {
	ports.CreateCandidateI
	ports.GetCandidateI
}

func CreateCandidatePort(implementation createCandidatePortTestRequires) {
	When("Given port is used", func() {
		testResource := domain.Candidate{Name: "Provided"}
		Context("When a valid candidate is provided", func() {
			It("Expect usecase to not raise an error and fetch candidate", func() {
				var err error
				var uuid string

				uuid, err = implementation.CreateCandidate(testResource)
				Expect(err).ToNot(HaveOccurred())
				Expect(uuid).ToNot(BeEmpty())
				candidate, err := implementation.GetCandidate(uuid)
				Expect(err).ToNot(HaveOccurred())
				Expect(candidate.Name).To(Equal(testResource.Name))
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
