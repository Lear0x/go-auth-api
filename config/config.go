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

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Avertissement: Impossible de charger le fichier .env, valeurs par défaut utilisées")
	}
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func ConnectDB() {
	var err error
	var err2 error

	dsn := GetEnv("DB_URL", "host=postgres user=admin password=pass dbname=authdb sslmode=disable")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Impossible de se connecter à la base de données :", err)
	} else {
		fmt.Println("✅ Connexion à la base de données réussie !")

		err = DB.AutoMigrate(&models.User{})
		err2 = DB.AutoMigrate(&models.BlacklistedToken{})
		if err != nil {
			log.Fatal("❌ Erreur lors de la migration :", err)
		} else {
			fmt.Println("✅ Migration terminée avec succès !")
		}

		if err2 != nil {
			log.Fatal("❌ Erreur lors de la migration :", err2)
		} else {
			fmt.Println("✅ Migration terminée avec succès !")
		}
	}
}
