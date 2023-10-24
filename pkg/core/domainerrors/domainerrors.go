package domainerrors

import (
	"fmt"
)

type VotingError string

func (e VotingError) Error() string {
	return string(e)
}

func (e VotingError) Unwrap() error {
	return error(e)
}

var ErrBaseVotingError = VotingError("Voting error")

var (
	ErrCandidateNotFound     = fmt.Errorf("%w: candidate not found", ErrBaseVotingError)
	ErrVotingSessionNotFound = fmt.Errorf("%w: voting session not found", ErrBaseVotingError)
)
