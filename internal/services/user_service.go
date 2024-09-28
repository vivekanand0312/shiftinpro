package service

import (
	"tg/internal/models"
	"tg/internal/repository"
)

// UserService defines the methods for user service operations
type UserService interface {
	GetUser(id uint) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetUser(id uint) (*models.User, error) {
	return s.repo.GetUserByID(id)
}
