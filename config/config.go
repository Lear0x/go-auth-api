package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Lear0x/go-auth-api/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variable globale pour la base de données
var DB *gorm.DB

// LoadEnv charge les variables d'environnement depuis un fichier .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Avertissement: Impossible de charger le fichier .env, valeurs par défaut utilisées")
	}
}

// GetEnv récupère une variable d'environnement avec une valeur par défaut
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// ConnectDB établit la connexion à la base de données PostgreSQL
func ConnectDB() {
	var err error

	// Récupérer l'URL de connexion depuis .env
	dsn := GetEnv("DB_URL", "host=postgres user=admin password=pass dbname=authdb sslmode=disable")

	// Initialiser la connexion avec GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Impossible de se connecter à la base de données :", err)
	} else {
		fmt.Println("✅ Connexion à la base de données réussie !")

		// Auto-migration pour créer la table users
		err = DB.AutoMigrate(&models.User{})
		if err != nil {
			log.Fatal("❌ Erreur lors de la migration :", err)
		} else {
			fmt.Println("✅ Migration terminée avec succès !")
		}
	}
}
