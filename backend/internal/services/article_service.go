package services

import (
	"github.com/daichi1002/book_app/backend/internal/models"
	"github.com/daichi1002/book_app/backend/internal/repositories"
)

type ArticleService struct {
	Repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) ArticleService {
	return ArticleService{Repo: repo}
}

func (s *ArticleService) GetArticles() ([]models.Article, error) {
	return s.Repo.GetArticles()
}
