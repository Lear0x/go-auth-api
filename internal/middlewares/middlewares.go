package middlewares

import (
	"net/http"
	"strings"

	"github.com/Lear0x/go-auth-api/config"
	"github.com/Lear0x/go-auth-api/internal/models"
	"github.com/Lear0x/go-auth-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token manquant"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		var blacklistedToken models.BlacklistedToken
		if err := config.DB.Where("token = ?", tokenString).First(&blacklistedToken).Error; err == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalide (blacklisté)"})
			c.Abort()
			return
		}

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalide"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
