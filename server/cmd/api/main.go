package main

import (
	"log"
	"net/http"

	"github.com/NicolaAve/noir/server/internal/handlers"
	"github.com/NicolaAve/noir/server/internal/middleware"
	"github.com/NicolaAve/noir/server/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.Connect()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "online",
			"message": "Noir API operativa e connessa al database PostgreSQL",
		})
	})

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	protected := router.Group("/api")
	protected.Use(middleware.RequireAuth())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID, exists := c.Get("user_id")
			if !exists {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossibile recuperare l'ID utente"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Accesso verificato",
				"user_id": userID,
			})
		})

		protected.POST("/cellars", handlers.CreateCellar)
		protected.POST("/wines", handlers.AddWine)
	}

	log.Println("Avvio server Noir sulla porta :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Errore critico: %v", err)
	}
}
