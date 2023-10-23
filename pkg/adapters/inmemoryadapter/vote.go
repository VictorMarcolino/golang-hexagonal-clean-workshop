package inmemoryadapter

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/utils"
)

var _ ports.GetVotesI = &InMemory{}
var _ ports.CreateVoteI = &InMemory{}

func (i *InMemory) GetVotes(votingUUID string) ([]domain.Vote, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	var votes []domain.Vote
	for _, vote := range i.voteStore {
		if vote.VotingUUID == votingUUID {
			votes = append(votes, vote)
		}
	}
	return votes, nil
}

func (i *InMemory) CreateVote(vote domain.Vote) (string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	uuid := utils.GenerateUUID()
	vote.UUID = uuid
	i.voteStore[vote.UUID] = vote

	return uuid, nil
}
