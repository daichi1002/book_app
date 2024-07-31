package services

import (
	"github.com/daichi1002/book_app/backend/internal/models"
	"github.com/daichi1002/book_app/backend/internal/repositories"
)

type UserService struct {
	Repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return UserService{Repo: repo}
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}
