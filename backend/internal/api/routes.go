package api

import (
	"database/sql"
	"net/http"

	"github.com/daichi1002/book_app/backend/internal/api/handlers"
	"github.com/daichi1002/book_app/backend/internal/repositories/mysql"
	"github.com/daichi1002/book_app/backend/internal/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) http.Handler {
	userRepo := mysql.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := &handlers.UserHandler{Service: userService}

	articleRepo := mysql.NewArticleRepository(db)
	articleService := services.NewArticleService(articleRepo)
	articleHandler := &handlers.ArticleHandler{Service: articleService}

	r := mux.NewRouter()
	r.HandleFunc("/user", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/articles", articleHandler.GetArticles).Methods("GET")
	return r
}
