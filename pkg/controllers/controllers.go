package controllers

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/controllers/candidatecontroller"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/controllers/votingcontroller"
)

var NewVotingController = votingcontroller.NewVotingController
var NewCandidateController = candidatecontroller.NewCandidateController
