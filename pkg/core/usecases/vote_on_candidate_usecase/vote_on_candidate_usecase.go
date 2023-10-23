package vote_on_candidate_usecase

import (
	"errors"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
)

type VoteOnCandidateUsecase struct {
	ports.GetCandidateI
	ports.GetVotingSessionI
	ports.CreateVoteI
}

func (v *VoteOnCandidateUsecase) ExecuteVoteOnCandidateUsecase(candidate domain.Candidate, voting domain.Voting) error {
	if candidate.UUID == "" {
		return errors.New("candidate UUID is missing, cannot vote without it")
	}
	if voting.UUID == "" {
		return errors.New("voting session UUID is missing, cannot vote without it")
	}

	// Check if the voting session exists
	_, err := v.GetVotingSession(voting.UUID)
	if err != nil {
		return errors.New("voting session does not exist")
	}

	// Check if the candidate is part of the voting session
	foundCandidate, err := v.GetCandidate(candidate.UUID)
	if err != nil || foundCandidate.UUID != candidate.UUID {
		return errors.New("candidate does not exist in the voting session")
	}

	// Record the vote
	vote := domain.Vote{
		UUID:          "", // UUID can be generated in the repository layer
		VotingUUID:    voting.UUID,
		CandidateUuid: candidate.UUID,
	}
	_, err = v.CreateVote(vote)
	if err != nil {
		return errors.New("failed to record the vote")
	}

	return nil
}

type requiredVoteOnCandidateUsecase interface {
	ports.GetCandidateI
	ports.GetVotingSessionI
	ports.CreateVoteI
}

func NewVoteOnCandidateUsecase(adapters requiredVoteOnCandidateUsecase) *VoteOnCandidateUsecase {
	return &VoteOnCandidateUsecase{
		GetCandidateI:     adapters,
		GetVotingSessionI: adapters,
		CreateVoteI:       adapters,
	}
}
