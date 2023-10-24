package count_votes_usecase

import (
	"errors"

	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
)

type GetVotingResultUsecase struct {
	ports.GetVotesI
	ports.GetVotingSessionI
}

func (g *GetVotingResultUsecase) ExecuteGetVotingResultUsecase(uuid string) (map[string]int, error) {
	if uuid == "" {
		return nil, errors.New("voting session UUID is missing")
	}
	_, err := g.GetVotingSession(uuid)
	if err != nil {
		return nil, errors.New("voting session does not exist")
	}
	// Check if the voting session exists
	votes, err := g.GetVotes(uuid)
	if err != nil {
		return nil, errors.New("error on votes")
	}

	// Count the votes for each candidate
	voteCounts := make(map[string]int)
	for _, vote := range votes {
		voteCounts[vote.CandidateUuid]++
	}

	return voteCounts, nil
}

type requiredInterface interface {
	ports.GetVotesI
	ports.GetVotingSessionI
}

func NewGetVotingResultUsecase(adapter requiredInterface) *GetVotingResultUsecase {
	return &GetVotingResultUsecase{
		GetVotesI:         adapter,
		GetVotingSessionI: adapter,
	}
}
