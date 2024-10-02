package services

import (
	"shiftinpro/internal/models"
	"shiftinpro/internal/repositories"
)

type BookingService interface {
	GetItemChecklists() ([]models.ItemChecklist, error)
}

type bookingService struct {
	bookingRepo repositories.BookingRepository
}

// NewBookingService creates a new instance of bookingService.
func NewBookingService(bookingRepo repositories.BookingRepository) BookingService {
	return &bookingService{bookingRepo: bookingRepo}
}

// GetItemChecklists fetches item checklists by calling the repository.
func (s *bookingService) GetItemChecklists() ([]models.ItemChecklist, error) {
	return s.bookingRepo.GetItemChecklists()
}
