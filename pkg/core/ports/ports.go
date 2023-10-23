package ports

import "github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"

type CreateCandidateI interface {
	CreateCandidate(candidate domain.Candidate) (uuid string, err error)
}
type CreateVotingI interface {
	CreateVoting(voting domain.Voting) (uuid string, err error)
}
type CreateVoteI interface {
	CreateVote(vote domain.Vote) (uuid string, err error)
}

type GetCandidateI interface {
	GetCandidate(uuid string) (candidate domain.Candidate, err error)
}
type GetVotingSessionI interface {
	GetVotingSession(uuid string) (voting domain.Voting, err error)
}
type GetVotesI interface {
	GetVotes(VotingUuid string) (votes []domain.Vote, err error)
}
