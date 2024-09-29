package models

import (
    "time"
)

type User struct {
    ID              uint    `gorm:"primaryKey"`
    Phone           string  `gorm:"unique;not null"`
    Name            string  `gorm:"not null"`
    Image           *string `gorm:"null"` // Path to the uploaded image
    UserType        int     `gorm:"not null"`
    IsOtpVerified   bool    `gorm:"default:false"`
    GovIdType       *string `gorm:"type:enum('aadhar', 'gst');default:null"`
    GovIdVal        *int    `gorm:"null"`
    IsGovIdVerified bool
    CreatedAt       time.Time `gorm:"autoCreateTime"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
