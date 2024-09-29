package services

import (
    "errors"

    "shiftproin/internal/models"
    "shiftproin/internal/repository"
)

type UserService interface {
    SaveUser(user *models.User) error
    GetUser(phone string) (*models.User, error)
    ValidateOTP(phone string, otp int) (bool, error)
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) SaveUser(user *models.User) error {
    return s.repo.CreateUser(user)
}

func (s *userService) GetUser(phone string) (*models.User, error) {
    return s.repo.GetUserByPhone(phone)
}

func (s *userService) ValidateOTP(phone string, otp int) (bool, error) {
    if phone == "" {
        return false, errors.New("invalid Phone")
    }

    if otp != 1234 {
        return false, errors.New("invalid OTP")
    }
    return true, nil
}
