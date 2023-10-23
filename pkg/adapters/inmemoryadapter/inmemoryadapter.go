package inmemoryadapter

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"sync"
)

type InMemory struct {
	candidateStore map[string]domain.Candidate
	voteStore      map[string]domain.Vote
	votingStore    map[string]domain.Voting
	mu             sync.RWMutex
}

func NewInMemoryAdapter() *InMemory {
	return &InMemory{
		candidateStore: make(map[string]domain.Candidate),
		voteStore:      make(map[string]domain.Vote),
		votingStore:    make(map[string]domain.Voting),
	}
}
