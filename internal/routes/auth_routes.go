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

	// 	api := router.Group("/19ebe88a-e0ce-42bc-8dcf-d5206d0658ad")

	// 	api.GET("/health", controllers.HealthCheck)
	// 	api.POST("/register", controllers.Register)
	// 	api.POST("/login", controllers.Login)
	// 	api.POST("/forgot-password", controllers.ForgotPassword)
	// 	api.POST("/reset-password", controllers.ResetPassword)
}
