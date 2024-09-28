package repository

import (
	"errors"
	"tg/internal/models"
)

// UserRepository defines the methods for user repository operations
type UserRepository interface {
	GetUserByID(id uint) (*models.User, error)
}

type userRepository struct {
	users map[uint]*models.User
}

func NewUserRepository() UserRepository {
	return &userRepository{
		users: map[uint]*models.User{
			1: {ID: 1, Name: "John Doe", Email: "john@example.com"},
		},
	}
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
