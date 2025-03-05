package main

import (
	"log"
	"net/http"

	"github.com/Lear0x/go-auth-api/config"
	"github.com/Lear0x/go-auth-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Charger les variables d'environnement
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Fichier .env non trouvé, chargement ignoré")
	}

	// Initialisation du routeur Gin
	router := gin.Default()

	// Ajouter les routes
	routes.SetupRoutes(router)

	// Récupérer le port depuis .env (par défaut : 8080)
	port := config.GetEnv("PORT", "8080")

	// Lancer le serveur
	log.Printf("🚀 Serveur démarré sur http://localhost:%s\n", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Erreur au démarrage du serveur :", err)
	}
}
