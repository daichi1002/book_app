package api

import (
	"database/sql"

	"github.com/daichi1002/book_app/backend/internal/api/handlers"
	"github.com/daichi1002/book_app/backend/internal/repositories/db"
	"github.com/daichi1002/book_app/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, d *sql.DB) {
	userRepo := db.NewUserRepository(d)
	userService := services.NewUserService(userRepo)
	userHandler := &handlers.UserHandler{Service: userService}

	articleRepo := db.NewArticleRepository(d)
	articleService := services.NewArticleService(articleRepo)
	articleHandler := &handlers.ArticleHandler{Service: articleService}

	e.GET("/user", userHandler.GetUser)
	e.GET("/articles", articleHandler.GetArticles)
	e.GET("/articles/:id", articleHandler.GetArticle)
	e.POST("/articles/create", articleHandler.CreateArticle)
}
