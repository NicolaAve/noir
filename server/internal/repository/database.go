package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/NicolaAve/noir/server/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Avviso: .env non trovato, uso variabili di sistema")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Impossibile connettersi al database: ", err)
	}

	err = database.AutoMigrate(&models.Wine{}, &models.Cellar{}, &models.User{}, &models.TastingHistory{})
	if err != nil {
		log.Println("Errore migrazione:", err)
	}

	DB = database
	log.Println("Connessione PostgreSQL stabilita e tabelle migrate.")
}
