package services

import (
	"shiftinpro/internal/models"
	"shiftinpro/internal/repositories"
)

type AddressService interface {
	FetchAddress(pincode *float64, state *string) ([]models.Address, error)
}

type addressService struct {
	repo repositories.AddressRepository
}

func NewAddressService(repo repositories.AddressRepository) AddressService {
	return &addressService{repo: repo}
}

func (s *addressService) FetchAddress(pincode *float64, state *string) ([]models.Address, error) {
	if pincode != nil {
		return s.repo.GetAddressesByPincode(*pincode)
	}
	if state != nil {
		return s.repo.GetAddressesByState(*state)
	}
	return nil, nil
}
