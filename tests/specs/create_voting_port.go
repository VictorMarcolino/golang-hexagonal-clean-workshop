package specs

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type createVotingPortTestRequires interface {
	ports.CreateVotingI
}

func CreateVotingPort(implementation createVotingPortTestRequires) {
	When("Given port is used", func() {
		var uuid string
		testResource := domain.Voting{Name: "Provided"}

		Context("When a valid voting session is provided", func() {
			It("Expect usecase to not raise an error", func() {
				var err error
				uuid, err = implementation.CreateVoting(testResource)
				Expect(err).ToNot(HaveOccurred())
				Expect(uuid).ToNot(BeEmpty())
			})
		})
	})
}
