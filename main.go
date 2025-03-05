package main

import (
	"log"
	"net/http"

	"github.com/Lear0x/go-auth-api/config"
	"github.com/Lear0x/go-auth-api/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Fichier .env non trouvé, chargement ignoré")
	}

	config.ConnectDB()

	router := gin.Default()

	routes.SetupRoutes(router)

	port := config.GetEnv("PORT", "8080")

	// Lancer le serveur
	log.Printf("🚀 Serveur démarré sur http://localhost:%s\n", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Erreur au démarrage du serveur :", err)
	}
}
