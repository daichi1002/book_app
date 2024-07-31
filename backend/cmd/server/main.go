package main

import (
	"log"
	"net/http"

	"github.com/daichi1002/book_app/backend/internal/api"
	"github.com/daichi1002/book_app/backend/internal/config"
	"github.com/daichi1002/book_app/backend/pkg/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	db, err := database.NewMySQL(cfg.DB.DSN)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	router := api.NewRouter(db)
	log.Printf("Server starting on %s", cfg.Server.Address)
	log.Fatal(http.ListenAndServe(cfg.Server.Address, router))
}
