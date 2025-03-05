package controllers

import (
	"fmt"
	"net/http"

	"github.com/Lear0x/go-auth-api/config"
	"github.com/Lear0x/go-auth-api/internal/models"
	"github.com/Lear0x/go-auth-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
		return
	}

	// Vérifier si l'utilisateur existe déjà
	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Cet email est déjà utilisé"})
		return
	}

	// Hasher le mot de passe avant de l'enregistrer
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de hachage du mot de passe"})
		return
	}

	fmt.Println("Mot de passe APRÈS hachage :", user.Password)

	// Enregistrer l'utilisateur en base de données
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur d'enregistrement"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Utilisateur créé"})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("🔴 Erreur binding JSON :", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
		return
	}

	fmt.Println("📥 Email reçu :", input.Email)
	fmt.Println("📥 Mot de passe reçu :", input.Password)

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	fmt.Println("🔍 Utilisateur trouvé :", user.Email)
	fmt.Println("🔍 Mot de passe récupéré en base :", user.Password)

	if !user.CheckPassword(input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Mot de passe incorrect"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de génération du token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
