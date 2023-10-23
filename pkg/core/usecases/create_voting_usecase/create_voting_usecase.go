package create_voting_usecase

import (
	"errors"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
)

type CreateVotingUsecase struct {
	ports.CreateVotingI
	ports.GetCandidateI
}

func (c *CreateVotingUsecase) ExecuteCreateVotingUsecase(name string, candidates []domain.Candidate) (string, error) {
	if name == "" {
		return "", errors.New("voting name is missing, cannot create a voting session without it")
	}
	if len(candidates) == 0 {
		return "", errors.New("candidate list is empty, cannot create a voting session without candidates")
	}

	votingSession := domain.Voting{
		Name:       name,
		Candidates: []domain.Candidate{},
	}

	for _, item := range candidates {
		if item.UUID == "" {
			return "", errors.New("candidate UUID is missing")
		}

		candidate, err := c.GetCandidate(item.UUID)
		if err != nil {
			return "", err
		}
		votingSession.Candidates = append(votingSession.Candidates, candidate)
	}

	votingUUID, err := c.CreateVoting(votingSession)
	if err != nil {
		return "", err
	}

	return votingUUID, nil
}

type requiredInterface interface {
	ports.CreateVotingI
	ports.GetCandidateI
}

func NewCreateVotingUsecase(adapters requiredInterface) *CreateVotingUsecase {
	return &CreateVotingUsecase{
		CreateVotingI: adapters,
		GetCandidateI: adapters,
	}
}
