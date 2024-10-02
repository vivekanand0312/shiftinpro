package repositories

import (
	"gorm.io/gorm"
	"log"
	"shiftinpro/internal/models"
)

type BookingRepository interface {
	GetItemChecklists() ([]models.ItemChecklist, error)
}

type bookingRepository struct {
	db *gorm.DB
}

// NewBookingRepository creates a new instance of bookingRepository with GORM.
func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

// GetItemChecklists fetches the item checklists from the database using GORM.
func (r *bookingRepository) GetItemChecklists() ([]models.ItemChecklist, error) {
	var checklists []models.ItemChecklist
	err := r.db.Table("sd_item_checklist").
		Select("item_id, storage_kind_id, category, display_name, len as length, width, height").
		Find(&checklists).Error
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, err
	}

	// Calculate area and volume for each item
	for i, item := range checklists {
		checklists[i].Area = item.Length * item.Width
		checklists[i].Volume = item.Length * item.Width * item.Height
	}

	return checklists, nil
}
