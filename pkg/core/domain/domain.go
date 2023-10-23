package domain

type Candidate struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type Voting struct {
	UUID       string      `json:"uuid"`
	Name       string      `json:"name"`
	Candidates []Candidate `json:"candidates"`
}

type Vote struct {
	UUID          string `json:"uuid"`
	VotingUUID    string `json:"votingUuid"`
	CandidateUuid string `json:"candidateUuid"`
}
