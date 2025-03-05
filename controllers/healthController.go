package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck vérifie que l'API fonctionne
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API is running"})
}
