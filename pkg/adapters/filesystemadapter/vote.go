package filesystemadapter

import (
	"encoding/json"
	"fmt"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/ports"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/utils"
	"strings"
)

var _ ports.GetVotesI = &FilesystemAdapter{}
var _ ports.CreateVoteI = &FilesystemAdapter{}

func getFsUuidForVote(uuid string) string {
	return fmt.Sprintf("vote/%v.json", uuid)
}

func filterVotesString(list []string) []string {
	var votesList []string
	for _, v := range list {
		if strings.HasPrefix(v, "vote/") {
			votesList = append(votesList, v)
		}
	}
	return votesList
}
func (i *FilesystemAdapter) GetVotes(votingUUID string) ([]domain.Vote, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	var votes []domain.Vote
	files, err := i.ArtifactRepo.ListArtifacts()
	if err != nil {
		return nil, err
	}

	filterVotes := filterVotesString(files)
	for _, file := range filterVotes {
		reader, err := i.ArtifactRepo.GetArtifact(file)
		if err != nil {
			return nil, err
		}
		defer reader.Close()

		var vote domain.Vote
		if err := json.NewDecoder(reader).Decode(&vote); err != nil {
			return nil, err
		}
		if vote.VotingUUID == votingUUID {
			votes = append(votes, vote)
		}
	}
	return votes, nil
}

func (i *FilesystemAdapter) CreateVote(vote domain.Vote) (string, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	UUID := utils.GenerateUUID()
	vote.UUID = UUID

	data, err := json.Marshal(vote)
	if err != nil {
		return "", err
	}
	err = i.ArtifactRepo.PushArtifact(getFsUuidForVote(UUID), strings.NewReader(string(data)))
	if err != nil {
		return "", err
	}
	return UUID, nil
}
