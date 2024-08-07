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

func (s *ArticleService) GetArticle(id int) (*models.Article, error) {
	return s.Repo.GetArticle(id)
}

func (s *ArticleService) CreateArticle(params models.Article) error {
	return s.Repo.CreateArticle(params)
}

func (s *ArticleService) DeleteArticle(id int) error {
	return s.Repo.DeleteArticle(id)
}

func (s *ArticleService) UpdateArticle(params models.Article) error {
	return s.Repo.UpdateArticle(params)
}
