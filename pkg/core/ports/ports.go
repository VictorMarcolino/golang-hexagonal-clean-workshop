package ports

import "github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"

type CreateCandidateI interface {
	CreateCandidate(candidate domain.Candidate) (uuid string, err error)
}
type GetCandidateI interface {
	GetCandidate(uuid string) (candidate domain.Candidate, err error)
}
