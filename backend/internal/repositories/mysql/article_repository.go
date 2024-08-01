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

func (r *ArticleRepository) GetArticle(id int) (*models.Article, error) {
	var article models.Article

	row, err := r.DB.Query("SELECT id, title, content, created_at FROM articles WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	if !row.Next() {
		return nil, nil
	}

	err = row.Scan(
		&article.ID,
		&article.Title,
		&article.Content,
		&article.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &article, nil
}
