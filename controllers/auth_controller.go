package controllers

import (
	"net/http"

	"github.com/Lear0x/go-auth-api/config"
	"github.com/Lear0x/go-auth-api/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
		return
	}

	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de hachage du mot de passe"})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur d'enregistrement"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Utilisateur créé"})
}

// func Login(c *gin.Context) {
// 	var input models.User
// 	var user models.User

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
// 		return
// 	}

// 	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur non trouvé"})
// 		return
// 	}

// 	if !user.CheckPassword(input.Password) {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Mot de passe incorrect"})
// 		return
// 	}

// 	token, err := utils.GenerateToken(user.ID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de génération du token"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"token": token})
// }
