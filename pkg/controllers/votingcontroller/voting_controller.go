package votingcontroller

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases"
	"github.com/gin-gonic/gin"
)

type VotingController struct {
	usecases.CreateVotingUsecaseI
	usecases.VoteOnCandidateUsecaseI
	usecases.GetVotingResultUsecaseI
}

func NewVotingController(
	createVotingUsecase usecases.CreateVotingUsecaseI,
	voteUsecase usecases.VoteOnCandidateUsecaseI,
	getVotingResultUsecase usecases.GetVotingResultUsecaseI,
) *VotingController {
	return &VotingController{
		createVotingUsecase,
		voteUsecase,
		getVotingResultUsecase,
	}
}

func (p *VotingController) RegisterRoutes(router gin.IRouter) {
	voting := router.Group("voting")
	voting.POST("", p.createVotingUsecaseHandler)
	voting.POST(":votingUuid/:candidateUuid/vote", p.voteOnCandidateHandler)
	voting.GET(":votingUuid/result", p.getVotingResultHandler) // Add this line
}

// createVotingUsecaseHandler
// @Summary Create a new voting Session
// @Description Create Voting Session with given name
// @ID create-voting-session
// @Accept  json
// @Produce  json
// @Param create-voting-request body votingcontroller.createVotingUsecaseHandler.request true "Create Voting Session"
// @Success 200 {object} votingcontroller.createVotingUsecaseHandler.response
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /voting [post]
func (p *VotingController) createVotingUsecaseHandler(c *gin.Context) {
	type request struct {
		Candidates  []domain.Candidate `json:"candidates"`
		SessionName string             `json:"sessionName"`
	}
	type response struct {
		SessionUUID string `json:"sessionUUID,omitempty"`
		Error       string `json:"error,omitempty"`
	}

	var req request
	// Bind the request body to our request struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	sessionUUID, err := p.ExecuteCreateVotingUsecase(req.SessionName, req.Candidates)

	var resp response
	if err != nil {
		resp.Error = err.Error()
		c.JSON(400, resp)
		return
	}
	resp.SessionUUID = sessionUUID
	c.JSON(200, resp)
}

// voteOnCandidateHandler
// @Summary Vote on a Candidate
// @Description Cast a vote for a candidate within a voting session
// @ID vote-on-candidate
// @Accept  json
// @Produce  json
// @Param votingUuid path string true "UUID of the Voting Session"
// @Param candidateUuid path string true "UUID of the Candidate"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /voting/{votingUuid}/{candidateUuid}/vote [post]
func (p *VotingController) voteOnCandidateHandler(c *gin.Context) {
	votingUuid := c.Param("votingUuid")
	candidateUuid := c.Param("candidateUuid")

	err := p.ExecuteVoteOnCandidateUsecase(domain.Candidate{UUID: candidateUuid}, domain.Voting{UUID: votingUuid})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Vote successfully cast"})
}

// getVotingResultHandler
// @Summary Get Voting Results
// @Description Retrieve the results of a voting session
// @ID get-voting-results
// @Produce  json
// @Param votingUuid path string true "UUID of the Voting Session"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /voting/{votingUuid}/result [get]
func (p *VotingController) getVotingResultHandler(c *gin.Context) {
	votingUuid := c.Param("votingUuid")

	results, err := p.ExecuteGetVotingResultUsecase(votingUuid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, results)
}
