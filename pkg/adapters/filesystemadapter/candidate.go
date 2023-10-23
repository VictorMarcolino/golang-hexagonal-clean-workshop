package filesystemadapter

import (
	"encoding/json"
	"fmt"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domainerrors"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/utils"
	"strings"
)

var _ ports.CreateCandidateI = &FilesystemAdapter{}
var _ ports.GetCandidateI = &FilesystemAdapter{}

func getFsUuidForCandidate(uuid string) string {
	return fmt.Sprintf("candidate/%v.json", uuid)
}
func (i *FilesystemAdapter) CreateCandidate(candidate domain.Candidate) (string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	UUID := utils.GenerateUUID()
	candidate.UUID = UUID

	data, err := json.Marshal(candidate)
	if err != nil {
		return "", err
	}
	err = i.ArtifactRepo.PushArtifact(getFsUuidForCandidate(UUID), strings.NewReader(string(data)))
	if err != nil {
		return "", err
	}
	return UUID, nil
}

func (i *FilesystemAdapter) GetCandidate(uuid string) (domain.Candidate, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	reader, err := i.ArtifactRepo.GetArtifact(getFsUuidForCandidate(uuid))
	if err != nil {
		return domain.Candidate{}, domainerrors.ErrCandidateNotFound
	}
	defer reader.Close()

	var candidate domain.Candidate
	if err := json.NewDecoder(reader).Decode(&candidate); err != nil {
		return domain.Candidate{}, err
	}
	return candidate, nil
}
