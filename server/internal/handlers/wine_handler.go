package handlers

import (
	"net/http"

	"github.com/NicolaAve/noir/server/internal/models"
	"github.com/NicolaAve/noir/server/internal/repository"
	"github.com/gin-gonic/gin"
)

func AddWine(c *gin.Context) {

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Accesso negato: utente non identificato"})
		return
	}

	uid := uint(userID.(float64))

	var input struct {
		Name     string  `json:"name" binding:"required"`
		CellarID uint    `json:"cellar_id" binding:"required"`
		Quantity int     `json:"quantity"`
		Producer string  `json:"producer"`
		Year     int     `json:"year"`
		Type     string  `json:"type"`
		Grape    string  `json:"grape"`
		Rating   int     `json:"rating"`
		Notes    string  `json:"notes"`
		ImageURL string  `json:"image_url"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dati non validi: 'name' e 'cellar_id' sono obbligatori"})
		return
	}

	var count int64
	repository.DB.Table("user_cellars").Where("user_id = ? AND cellar_id = ?", uid, input.CellarID).Count(&count)
	if count == 0 {
	
		c.JSON(http.StatusForbidden, gin.H{"error": "Azione non consentita: non hai accesso a questa cantina"})
		return
	}

	quantity := input.Quantity
	if quantity <= 0 {
		quantity = 1
	}

	wine := models.Wine{
		Name:     input.Name,
		CellarID: input.CellarID,
		Quantity: quantity,
		Status:   "in_stock", 
		Producer: input.Producer,
		Year:     input.Year,
		Type:     input.Type,
		Grape:    input.Grape,
		Rating:   input.Rating,
		Notes:    input.Notes,
		ImageURL: input.ImageURL,
	}

	if err := repository.DB.Create(&wine).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore durante il salvataggio del vino"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Vino aggiunto alla cantina con successo",
		"wine":    wine,
	})
}