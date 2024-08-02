package api

import (
	"database/sql"

	"github.com/daichi1002/book_app/backend/internal/api/handlers"
	"github.com/daichi1002/book_app/backend/internal/repositories/mysql"
	"github.com/daichi1002/book_app/backend/internal/services"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, db *sql.DB) {
	userRepo := mysql.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := &handlers.UserHandler{Service: userService}

	articleRepo := mysql.NewArticleRepository(db)
	articleService := services.NewArticleService(articleRepo)
	articleHandler := &handlers.ArticleHandler{Service: articleService}

	e.GET("/user", userHandler.GetUser)
	e.GET("/articles", articleHandler.GetArticles)
	e.GET("/articles/:id", articleHandler.GetArticle)
	e.POST("/articles/create", articleHandler.CreateArticle)
}
