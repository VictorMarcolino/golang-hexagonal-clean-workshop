package usecases

import "github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"

type CreateCandidateUsecaseI interface {
	ExecuteCreateCandidate(candidate domain.Candidate) (uuid string, err error)
}
