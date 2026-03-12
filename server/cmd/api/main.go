package main

import (
	"log"
	"net/http"

	"github.com/NicolaAve/noir/server/internal/handlers"
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

	log.Println("Avvio server Noir sulla porta :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Errore critico: %v", err)
	}
}