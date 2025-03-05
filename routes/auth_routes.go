package routes

import (
	"github.com/Lear0x/go-auth-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", controllers.HealthCheck)

	// router.POST("/register", controllers.Register)
	// router.POST("/login", controllers.Login)
}
