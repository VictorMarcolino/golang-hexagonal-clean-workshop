package inmemoryadapter

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domainerrors"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/utils"
)

var (
	_ ports.CreateCandidateI = &InMemory{}
	_ ports.GetCandidateI    = &InMemory{}
)

func (i *InMemory) CreateCandidate(candidate domain.Candidate) (string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	UUID := utils.GenerateUUID()
	candidate.UUID = UUID
	i.candidateStore[UUID] = candidate

	return UUID, nil
}

func (i *InMemory) GetCandidate(uuid string) (domain.Candidate, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	candidate, exists := i.candidateStore[uuid]
	if !exists {
		return domain.Candidate{}, domainerrors.ErrCandidateNotFound
	}

	return candidate, nil
}
