package main

import (
	"log"
	"net/http"
	"os"

	"github.com/daichi1002/book_app/backend/internal/api"
	"github.com/daichi1002/book_app/backend/internal/config"
	"github.com/daichi1002/book_app/backend/pkg/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	db, err := database.NewPostgreAQL(cfg.DB.DSN)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	defer db.Close()

	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("CORS_URL")}, // フロントエンドのオリジンを設定
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))
	api.SetupRoutes(e, db)

	e.Logger.Fatal(e.Start(cfg.Server.Address))
}
