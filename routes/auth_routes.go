package routes

import (
	"github.com/Lear0x/go-auth-api/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes ajoute toutes les routes de l'API
func SetupRoutes(router *gin.Engine) {
	// Route de santé
	router.GET("/health", controllers.HealthCheck)
}
