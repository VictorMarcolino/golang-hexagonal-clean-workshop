package usecases

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases/create_candidate_usecase"
)

type CreateCandidateUsecaseI interface {
	ExecuteCreateCandidate(candidate domain.Candidate) (uuid string, err error)
}

var NewCreateCandidateUsecase = create_candidate_usecase.NewCreateCandidateUsecase
