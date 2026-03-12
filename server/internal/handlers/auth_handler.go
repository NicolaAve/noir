package handlers

import (
	"net/http"

	"github.com/NicolaAve/noir/server/internal/models"
	"github.com/NicolaAve/noir/server/internal/repository"
	"github.com/NicolaAve/noir/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

// funzione per gestire la registrazione
func Register(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dati non validi: controlla email e password (minimo 6 caratteri)"})
		return
	}

	// criptazione password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore interno del server"})
		return
	}

	// creazione oggetto utente
	user := models.User{
		Email:    input.Email,
		Password: hashedPassword,
	}

	// salviamo, se esiste gia riceviamo errore
	if result := repository.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email già registrata"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Utente registrato con successo", "email": user.Email})
}