package controllers

import (
	model "Executor/models"
	services "Executor/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExecutorRequestController interface {
	ExecutorRequest(ctx *gin.Context)
}

type ExecutorRequestcontroller struct {
	repo services.ExecutorRepository
}

func NewExecutorRequestController(repo services.ExecutorRepository) ExecutorRequestController {
	return &ExecutorRequestcontroller{
		repo: repo,
	}
}

// ExecutorRequest handles the incoming HTTP request for executing an action.
// It binds the incoming JSON payload to a `model.Test` object, processes it via a repository call,
// and returns the result as a JSON response.
func (c *ExecutorRequestcontroller) ExecutorRequest(ctx *gin.Context) {
	rq := model.Test{}
	if err := ctx.BindJSON(&rq); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// Process the request by passing the 'Value' from the bound model to the repository function
	request := c.repo.ExecutorRequest(rq.Value)
	ctx.JSON(http.StatusOK, request)
}
