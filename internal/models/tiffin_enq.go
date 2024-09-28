package models

import (
	"fmt"
	"time"
)

// Enquiry represents the enquiries table in the database
type TiffinEnq struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name" binding:"required"`
	Mobile    string     `json:"mobile" db:"mobile" binding:"required"`
	Location  string     `json:"location" db:"location" binding:"required"`
	StartDate CustomDate `json:"start_date" db:"start_date" binding:"required"`
}

// CustomDate wraps time.Time to implement custom parsing
type CustomDate struct {
	time.Time
}

// UnmarshalJSON parses the date from a custom format
func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	// Remove quotes from the JSON string
	dateStr := string(b)
	dateStr = dateStr[1 : len(dateStr)-1]

	// Adjust the date format if month or day are not zero-padded
	parsedDate, err := parseCustomDate(dateStr)
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}

	cd.Time = parsedDate
	return nil
}

// MarshalJSON formats the date back to JSON in "YYYY-MM-DD"
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", cd.Time.Format("2006-01-02"))
	return []byte(formatted), nil
}

// parseCustomDate tries to parse date with various formats
func parseCustomDate(dateStr string) (time.Time, error) {
	formats := []string{
		"2006-01-02", // standard format
		"2006-1-2",   // non-zero-padded month/day
		"2006-01-2",  // zero-padded month
		"2006-1-02",  // zero-padded day
	}

	var parsedDate time.Time
	var err error

	for _, format := range formats {
		parsedDate, err = time.Parse(format, dateStr)
		if err == nil {
			return parsedDate, nil
		}
	}
	return time.Time{}, fmt.Errorf("could not parse date: %s", dateStr)
}
