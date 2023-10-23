package usecases

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases/count_votes_usecase"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases/create_candidate_usecase"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases/create_voting_usecase"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases/vote_on_candidate_usecase"
)

type CreateCandidateUsecaseI interface {
	ExecuteCreateCandidate(candidate domain.Candidate) (uuid string, err error)
}

var NewCreateCandidateUsecase = create_candidate_usecase.NewCreateCandidateUsecase

type VoteOnCandidateUsecaseI interface {
	ExecuteVoteOnCandidateUsecase(candidate domain.Candidate, voting domain.Voting) (err error)
}

var NewVoteOnCandidateUsecase = vote_on_candidate_usecase.NewVoteOnCandidateUsecase

type CreateVotingUsecaseI interface {
	ExecuteCreateVotingUsecase(name string, candidate []domain.Candidate) (uuid string, err error)
}

var NewCreateVotingUsecase = create_voting_usecase.NewCreateVotingUsecase

type GetVotingResultUsecaseI interface {
	ExecuteGetVotingResultUsecase(uuid string) (map[string]int, error)
}

var NewGetVotingResultUsecase = count_votes_usecase.NewGetVotingResultUsecase
