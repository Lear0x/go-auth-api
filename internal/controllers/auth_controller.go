package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Lear0x/go-auth-api/config"
	"github.com/Lear0x/go-auth-api/internal/models"
	"github.com/Lear0x/go-auth-api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func ForgotPassword(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email invalide"})
		return
	}

	// Vérifier si l'utilisateur existe
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	// Générer un token temporaire (expire en 30 minutes)
	expirationTime := time.Now().Add(30 * time.Minute).Unix()
	secretKey := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": expirationTime,
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la génération du token"})
		return
	}

	// 🔹 Normalement ici, on enverrait un email avec le lien contenant ce token.
	// Pour tester, on va juste afficher le token en réponse.
	resetLink := fmt.Sprintf("http://localhost:8080/reset-password?token=%s", tokenString)

	c.JSON(http.StatusOK, gin.H{
		"message":     "Un lien de réinitialisation a été envoyé",
		"reset_token": tokenString, // À remplacer par un vrai envoi d'email plus tard
		"reset_link":  resetLink,   // Juste pour test
	})
}

// ResetPassword vérifie le token et met à jour le mot de passe
// ResetPassword met à jour le mot de passe après vérification du token
func ResetPassword(c *gin.Context) {
	var input struct {
		Token       string `json:"token" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
		return
	}

	// Vérifier le token
	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(input.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Methode de signature invalide")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalide ou expiré"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token corrompu"})
		return
	}

	// Récupérer l'ID utilisateur depuis le token
	userID := uint(claims["sub"].(float64))

	// Récupérer l'utilisateur
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	// 🔍 Vérifier avant mise à jour
	fmt.Println("🔍 Ancien mot de passe en base :", user.Password)

	// Hacher le nouveau mot de passe
	user.Password = input.NewPassword
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du hachage du mot de passe"})
		return
	}

	// 🔍 Vérifier après hachage
	fmt.Println("🔐 Nouveau mot de passe haché :", user.Password)

	// Mettre à jour le mot de passe en base
	if err := config.DB.Model(&user).Update("password", user.Password).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de mettre à jour le mot de passe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mot de passe mis à jour avec succès"})
}
