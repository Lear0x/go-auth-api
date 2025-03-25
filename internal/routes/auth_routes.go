package routes

import (
	"github.com/Lear0x/go-auth-api/internal/controllers"

	"github.com/Lear0x/go-auth-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/19ebe88a-e0ce-42bc-8dcf-d5206d0658ad")
	{
		api.GET("/health", controllers.HealthCheck)
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.POST("/forgot-password", controllers.ForgotPassword)
		api.POST("/reset-password", controllers.ResetPassword)

		protected := api.Group("/user")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.GET("/me", controllers.Me)
			protected.POST("/logout", controllers.Logout)
		}
	}
}
