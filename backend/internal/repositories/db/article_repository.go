package db

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
	row, err := r.DB.Query("SELECT id, title, content, created_at FROM articles WHERE is_deleted = false;")
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

	row, err := r.DB.Query("SELECT id, title, content, created_at FROM articles WHERE id = $1 AND is_deleted = false;", id)
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

func (r *ArticleRepository) CreateArticle(params models.Article) error {

	_, err := r.DB.Exec(
		`INSERT INTO articles (
			title,
			content
		) VALUES ($1, $2);
		`,
		params.Title, params.Content,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *ArticleRepository) DeleteArticle(id int) error {

	_, err := r.DB.Exec(
		`
			UPDATE
				articles
			SET
				is_deleted = TRUE
			WHERE
				id = $1;
		`,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *ArticleRepository) UpdateArticle(params models.Article) error {

	_, err := r.DB.Exec(
		`
		UPDATE articles
		SET title = $1,
			content = $2
		WHERE
			id = $3
		`,
		params.Title, params.Content, params.ID,
	)

	if err != nil {
		return err
	}

	return nil
}
