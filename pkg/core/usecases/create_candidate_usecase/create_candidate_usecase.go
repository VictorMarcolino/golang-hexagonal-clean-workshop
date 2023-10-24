package create_candidate_usecase

import (
	"errors"

	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
)

type CreateCandidateUsecase struct {
	ports.CreateCandidateI
}

func (r *CreateCandidateUsecase) ExecuteCreateCandidate(candidate domain.Candidate) (string, error) {
	if candidate.Name == "" {
		return "", errors.New("name is missing, cannot create a candidate without it")
	}
	return r.CreateCandidate(candidate)
}

func NewCreateCandidateUsecase(requiredPorts ports.CreateCandidateI) *CreateCandidateUsecase {
	return &CreateCandidateUsecase{
		requiredPorts,
	}
}
