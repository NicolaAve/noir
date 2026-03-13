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
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=6"`
		Role      string `json:"role"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dati non validi: controlla email e password (minimo 6 caratteri)"})
		return
	}

	userRole := "user"
	if input.Role == "shop" {
		userRole = "shop"
	}


	// criptazione password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore interno del server"})
		return
	}

	// creazione oggetto utente
	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hashedPassword,
		Role:      userRole,
	}

	// salviamo, se esiste gia riceviamo errore
	if result := repository.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email già registrata"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Utente registrato con successo",
		"email": user.Email,
		"name":  user.FirstName,
		"role":  user.Role,

	})
}

// verifichiamo credernziali e rilasciamo il pass (jwt)
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dati mancanti o malformati"})
		return
	}

	var user models.User
	
	// 1. cerchiamo utente tramite mail (unique)
	if err := repository.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenziali non valide"})
		return
	}

	// confronto password con hash nel DB
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenziali non valide"})
		return
	}

	// geeriamo il token per il telefono
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore durante la generazione del token"})
		return
	}

	// consegnamo il token
	c.JSON(http.StatusOK, gin.H{
		"message": "Accesso consentito",
		"token":   token,
	})
}