package db

import (
	"database/sql"

	"github.com/daichi1002/book_app/backend/internal/models"
	"github.com/daichi1002/book_app/backend/internal/repositories"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	err := r.DB.QueryRow("SELECT id, name, age FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return nil, err
	}
	return user, nil
}
