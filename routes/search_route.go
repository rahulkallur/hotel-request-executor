package routes

import (
	"Executor/controllers"
	"Executor/services"

	"github.com/gin-gonic/gin"
)

func LoadExecutorRequestRoute(router *gin.Engine) {
	repo := services.NewExecutorRequestRepository()
	controller := controllers.NewExecutorRequestController(repo)
	onboarding := router.Group("Executor")
	onboarding.POST("ExecutorRequest", controller.ExecutorRequest)
}
