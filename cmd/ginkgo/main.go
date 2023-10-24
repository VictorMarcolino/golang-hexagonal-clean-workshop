package main

import (
	"log"

	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/cmd/ginkgo/docs"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/cmd/ginkgo/middleware"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/adapters"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/controllers"
	"github.com/VictorMarcolino/golang-hexagonal-clean-workshop/pkg/core/usecases"
	"github.com/gin-gonic/gin"
)

// @title Voting API
// @version 1.0
// @description This is a simple voting api.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	ginEngine := gin.Default()
	ginEngine.Use(middleware.CORSMiddleware())
	v1 := ginEngine.Group("api/v1")
	docs.RegisterSwagger(v1)

	artifactRepo := adapters.NewFileSystemAdapter()

	createVotingUsecase := usecases.NewCreateVotingUsecase(artifactRepo)
	voteOnCandidateUsecase := usecases.NewVoteOnCandidateUsecase(artifactRepo)
	getVotingResultUsecase := usecases.NewGetVotingResultUsecase(artifactRepo)
	createCandidateUsecase := usecases.NewCreateCandidateUsecase(artifactRepo)

	votingController := controllers.NewVotingController(createVotingUsecase, voteOnCandidateUsecase, getVotingResultUsecase)
	candidateController := controllers.NewCandidateController(createCandidateUsecase)
	votingController.RegisterRoutes(v1)
	candidateController.RegisterRoutes(v1)

	HOST := "0.0.0.0:8080"
	log.Printf("Swagger at: %v/api/v1/swagger/index.html", HOST)
	log.Fatal(ginEngine.Run(HOST))
}
