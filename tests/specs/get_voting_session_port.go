package specs

import (
	"errors"

	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domainerrors"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type getVotingSessionPortTestRequires interface {
	ports.GetVotingSessionI
	ports.CreateVotingI
}

func GetVotingSessionPort(implementation getVotingSessionPortTestRequires) {
	When("Given port is used", func() {
		var validUUID string
		expectedVotingSession := domain.Voting{Name: "Valid Voting Session"}

		BeforeEach(func() {
			var err error
			validUUID, err = implementation.CreateVoting(expectedVotingSession)
			Expect(err).ToNot(HaveOccurred())
		})

		Context("When a valid UUID is provided", func() {
			It("Should be able to fetch the corresponding voting session", func() {
				votingSession, err := implementation.GetVotingSession(validUUID)
				Expect(err).ToNot(HaveOccurred())
				expectedVotingSession.UUID = validUUID // Ensure the expected voting session has the right UUID
				Expect(votingSession).To(Equal(expectedVotingSession))
			})
		})

		Context("When a nonexistent UUID is provided", func() {
			nonexistentUUID := "12345-nonexistent"

			It("Should raise an error", func() {
				_, err := implementation.GetVotingSession(nonexistentUUID)
				Expect(err).To(HaveOccurred())
				Expect(errors.Is(err, domainerrors.ErrVotingSessionNotFound)).To(BeTrue())
			})
		})
	})
}
