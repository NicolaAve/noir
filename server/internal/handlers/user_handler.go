package handlers

import (
	"net/http"

	"github.com/NicolaAve/noir/server/internal/models"
	"github.com/NicolaAve/noir/server/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Accesso negato: utente non identificato"})
		return
	}

	uidFloat, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore interno: formato token corrotto"})
		return
	}
	uid := uint(uidFloat)

	var user models.User

	if err := repository.DB.Select("id", "first_name", "last_name", "email", "role").First(&user, uid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utente non trovato"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"role":       user.Role,
	})
}
