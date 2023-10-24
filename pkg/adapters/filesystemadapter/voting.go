package filesystemadapter

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domainerrors"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/utils"
)

var (
	_ ports.GetVotingSessionI = &FilesystemAdapter{}
	_ ports.CreateVotingI     = &FilesystemAdapter{}
)

func getFsUuidForVoting(uuid string) string {
	return fmt.Sprintf("voting/%v.json", uuid)
}

func (i *FilesystemAdapter) CreateVoting(voting domain.Voting) (string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	UUID := utils.GenerateUUID()
	voting.UUID = UUID

	data, err := json.Marshal(voting)
	if err != nil {
		return "", err
	}
	err = i.ArtifactRepo.PushArtifact(getFsUuidForVoting(UUID), strings.NewReader(string(data)))
	if err != nil {
		return "", err
	}
	return UUID, nil
}

func (i *FilesystemAdapter) GetVotingSession(uuid string) (domain.Voting, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	reader, err := i.ArtifactRepo.GetArtifact(getFsUuidForVoting(uuid))
	if err != nil {
		return domain.Voting{}, domainerrors.ErrVotingSessionNotFound
	}
	defer reader.Close()

	var voting domain.Voting
	if err := json.NewDecoder(reader).Decode(&voting); err != nil {
		return domain.Voting{}, err
	}
	return voting, nil
}
