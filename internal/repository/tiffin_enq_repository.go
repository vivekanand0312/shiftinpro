package repository

import (
	"database/sql"
	"errors"
	"tg/internal/models"
	"time"
)

// TiffinEnqRepository defines the methods for user repository operations
type TiffinEnqRepository interface {
	SaveTiffinEnq(enquiry models.TiffinEnq) (int, error)
	GetTiffinEnqByID(id int) (*models.TiffinEnq, error)
}

// tiffinEnqRepository struct for working with the DB
type tiffinEnqRepository struct {
	db *sql.DB
}

// NewTiffinEnqRepository creates a new instance of TiffinEnqRepository
func NewTiffinEnqRepository(db *sql.DB) TiffinEnqRepository {
	return &tiffinEnqRepository{
		db: db,
	}
}

// SaveTiffinEnq saves an enquiry to the database
func (r *tiffinEnqRepository) SaveTiffinEnq(enquiry models.TiffinEnq) (int, error) {
	if enquiry.Name == "" || enquiry.Mobile == "" || enquiry.Location == "" || enquiry.StartDate.Before(time.Now()) {
		return 0, errors.New("validation failed: check required fields or invalid start date")
	}

	query := "INSERT INTO enquiries (name, mobile, location, start_date) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, enquiry.Name, enquiry.Mobile, enquiry.Location, enquiry.StartDate)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// GetTiffinEnqByID retrieves an enquiry by its ID
func (r *tiffinEnqRepository) GetTiffinEnqByID(id int) (*models.TiffinEnq, error) {
	var enquiry models.TiffinEnq
	query := "SELECT id, name, mobile, location, start_date FROM enquiries WHERE id = ?"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&enquiry.ID, &enquiry.Name, &enquiry.Mobile, &enquiry.Location, &enquiry.StartDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("enquiry not found")
		}
		return nil, err
	}

	return &enquiry, nil
}
