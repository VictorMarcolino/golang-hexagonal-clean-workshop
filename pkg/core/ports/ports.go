package ports

import "github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"

type CreateCandidateI interface {
	CreateCandidate(candidate domain.Candidate) (string, error)
}
