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

		// Extraire le token (supprimer "Bearer ")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Vérifier si le token est dans la blacklist
		var blacklistedToken models.BlacklistedToken
		if err := config.DB.Where("token = ?", tokenString).First(&blacklistedToken).Error; err == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalide (blacklisté)"})
			c.Abort()
			return
		}

		// Vérifier et décoder le token avec `utils.VerifyToken()`
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalide"})
			c.Abort()
			return
		}

		// Ajouter l'ID utilisateur dans le contexte Gin
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
