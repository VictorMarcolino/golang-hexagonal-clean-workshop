package inmemoryadapter

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domainerrors"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/utils"
)

var (
	_ ports.GetVotingSessionI = &InMemory{}
	_ ports.CreateVotingI     = &InMemory{}
)

func (i *InMemory) CreateVoting(voting domain.Voting) (string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	uuid := utils.GenerateUUID()
	voting.UUID = uuid
	i.votingStore[uuid] = voting

	return uuid, nil
}

func (i *InMemory) GetVotingSession(uuid string) (domain.Voting, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	voting, exists := i.votingStore[uuid]
	if !exists {
		return domain.Voting{}, domainerrors.ErrVotingSessionNotFound
	}

	return voting, nil
}
