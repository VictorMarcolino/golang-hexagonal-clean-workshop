package candidatecontroller

import (
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/domain"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases"
	"github.com/gin-gonic/gin"
)

type CandidateController struct {
	usecases.CreateCandidateUsecaseI
}

func NewCandidateController(
	createCandidate usecases.CreateCandidateUsecaseI,
) *CandidateController {
	return &CandidateController{
		createCandidate,
	}
}

func (p *CandidateController) RegisterRoutes(router gin.IRouter) {
	candidate := router.Group("candidate")
	candidate.POST("", p.createCandidateUsecaseHandler)
}

type createCandidateRequest struct {
	Name string `json:"name"`
}

type createCandidateResponse struct {
	UUID  string `json:"uuid,omitempty"`
	Error string `json:"error,omitempty"`
}

// createCandidateUsecaseHandler
// @Summary Create a new candidate
// @Description Create a new candidate with a given name
// @ID create-candidate
// @Accept  json
// @Produce  json
// @Param candidate body createCandidateRequest true "Candidate Information"
// @Success 200 {object} createCandidateResponse
// @Failure 400 {object} createCandidateResponse
// @Failure 500 {object} createCandidateResponse
// @Router /candidate [post]
func (p *CandidateController) createCandidateUsecaseHandler(c *gin.Context) {
	var req createCandidateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, createCandidateResponse{Error: "Invalid request format"})
		return
	}

	uuid, err := p.ExecuteCreateCandidate(domain.Candidate{Name: req.Name})
	if err != nil {
		c.JSON(400, createCandidateResponse{Error: err.Error()})
		return
	}

	c.JSON(200, createCandidateResponse{UUID: uuid})
}
