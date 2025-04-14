package routes

import (
	"github.com/DAT-CANDIDATE/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the API routes
func SetupRoutes(router *gin.Engine) {
	router.POST("/candidates", controllers.CreateCandidate)
	router.GET("/candidates", controllers.GetCandidates)
	router.GET("/candidates/:unique_id", controllers.GetCandidateByUniqueId)
	router.PUT("/candidates/:unique_id", controllers.UpdateCandidate)
}
