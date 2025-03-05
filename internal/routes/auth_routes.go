package routes

import (
	"github.com/Lear0x/go-auth-api/internal/controllers"

	"github.com/Lear0x/go-auth-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", controllers.HealthCheck)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/forgot-password", controllers.ForgotPassword)
	router.POST("/reset-password", controllers.ResetPassword)

	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	protected.POST("/logout", controllers.Logout)
	protected.GET("/me", controllers.Me)
}
