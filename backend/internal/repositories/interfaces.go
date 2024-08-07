package repositories

import "github.com/daichi1002/book_app/backend/internal/models"

type UserRepository interface {
	GetUserByID(id string) (*models.User, error)
}

type ArticleRepository interface {
	GetArticles() ([]models.Article, error)
	GetArticle(id int) (*models.Article, error)
	CreateArticle(params models.Article) error
	DeleteArticle(id int) error
	UpdateArticle(params models.Article) error
}
