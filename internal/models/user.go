package models

import (
    "time"
)

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Phone     string    `gorm:"unique;not null"`
    Name      string    `gorm:"not null"`
    Image     *string   `gorm:"null"` // Path to the uploaded image
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
