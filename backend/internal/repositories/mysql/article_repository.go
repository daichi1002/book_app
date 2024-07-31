package mysql

import (
	"database/sql"

	"github.com/daichi1002/book_app/backend/internal/models"
	"github.com/daichi1002/book_app/backend/internal/repositories"
)

type ArticleRepository struct {
	DB *sql.DB
}

func NewArticleRepository(db *sql.DB) repositories.ArticleRepository {
	return &ArticleRepository{DB: db}
}

func (r *ArticleRepository) GetArticles() ([]models.Article, error) {
	var articles []models.Article
	row, err := r.DB.Query("SELECT id, title, content, created_at FROM articles")
	if err != nil {
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		var elem models.Article

		err := row.Scan(
			&elem.ID,
			&elem.Title,
			&elem.Content,
			&elem.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		articles = append(articles, elem)
	}
	return articles, nil
}
