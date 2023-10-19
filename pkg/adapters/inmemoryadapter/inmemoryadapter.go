package inmemoryadapter

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"sync"
)

type InMemory struct {
	candidateStore map[string]domain.Candidate
	mu             sync.RWMutex
}

func NewInMemoryAdapter() *InMemory {
	return &InMemory{
		candidateStore: make(map[string]domain.Candidate),
	}
}
