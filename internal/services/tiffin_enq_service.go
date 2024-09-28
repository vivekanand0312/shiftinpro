package service

import (
	"tg/internal/models"
	"tg/internal/repository"
)

// TiffinEnqService interface for tiffinEnq service
type TiffinEnqService interface {
	CreateTiffinEnq(tiffinEnq models.TiffinEnq) (int, error)
	GetTiffinEnq(id int) (*models.TiffinEnq, error)
}

// tiffinEnqService struct to implement TiffinEnqService interface
type tiffinEnqService struct {
	repo repository.TiffinEnqRepository
}

// NewTiffinEnqService creates a new instance of TiffinEnqService
func NewTiffinEnqService(repo repository.TiffinEnqRepository) TiffinEnqService {
	return &tiffinEnqService{
		repo: repo,
	}
}

// CreateTiffinEnq creates a new tiffinEnq
func (s *tiffinEnqService) CreateTiffinEnq(tiffinEnq models.TiffinEnq) (int, error) {
	return s.repo.SaveTiffinEnq(tiffinEnq)
}

// GetTiffinEnq retrieves an tiffinEnq by its ID
func (s *tiffinEnqService) GetTiffinEnq(id int) (*models.TiffinEnq, error) {
	return s.repo.GetTiffinEnqByID(id)
}
