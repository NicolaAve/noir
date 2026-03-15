package handlers

import (
	"net/http"

	"github.com/NicolaAve/noir/server/internal/models"
	"github.com/NicolaAve/noir/server/internal/repository"
	"github.com/gin-gonic/gin"
)

func CreateCellar(c *gin.Context) {

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

	var input struct {
		Name     string `json:"name" binding:"required"`
		Location string `json:"location"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dati non validi: il nome della cantina è obbligatorio"})
		return
	}

	cellar := models.Cellar{
		Name:     input.Name,
		Location: input.Location,
	}

	if err := repository.DB.Create(&cellar).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore durante la creazione della cantina"})
		return
	}

	var user models.User

	if err := repository.DB.First(&user, uid).Error; err == nil {
		repository.DB.Model(&user).Association("Cellars").Append(&cellar)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Cantina creata con successo",
		"cellar":  cellar,
	})
}

func GetMyCellars(c *gin.Context) {

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

	if err := repository.DB.Preload("Cellars").First(&user, uid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore durante il recupero delle cantine"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cellars": user.Cellars,
	})
}
