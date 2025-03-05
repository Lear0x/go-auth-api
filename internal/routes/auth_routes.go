package routes

import (
	"github.com/Lear0x/go-auth-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", controllers.HealthCheck)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/forgot-password", controllers.ForgotPassword)
	router.POST("/reset-password", controllers.ResetPassword)
}
