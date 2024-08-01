package main

import (
	"log"

	"github.com/daichi1002/book_app/backend/internal/api"
	"github.com/daichi1002/book_app/backend/internal/config"
	"github.com/daichi1002/book_app/backend/pkg/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	defer db.Close()

	e := echo.New()
	api.SetupRoutes(e, db)

	log.Fatal(e.Start(cfg.Server.Address))
}
