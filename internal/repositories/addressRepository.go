package repositories

import (
	"gorm.io/gorm"
	"shiftinpro/internal/models"
)

type AddressRepository interface {
	GetAddressesByPincode(pincode float64) ([]models.Address, error)
	GetAddressesByState(state string) ([]models.Address, error)
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db: db}
}

func (r *addressRepository) GetAddressesByPincode(pincode float64) ([]models.Address, error) {
	var addresses []models.Address
	result := r.db.Table("sd_address").
		Select("sd_address.*, sd_countries.id as country_id, sd_countries.sortname, sd_countries.name").
		Joins("JOIN sd_countries ON sd_countries.id = sd_address.country_id").
		Where("sd_address.pincode = ?", pincode).
		Scan(&addresses)
	return addresses, result.Error
}

func (r *addressRepository) GetAddressesByState(state string) ([]models.Address, error) {
	var addresses []models.Address
	result := r.db.Table("sd_address").
		Select("sd_address.*, sd_countries.id as country_id, sd_countries.sortname, sd_countries.name").
		Joins("JOIN sd_countries ON sd_countries.id = sd_address.country_id").
		Where("sd_address.state = ?", state).
		Scan(&addresses)
	return addresses, result.Error
}
