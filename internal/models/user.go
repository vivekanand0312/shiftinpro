package models

import (
    "time"
)

type User struct {
    ID              uint      `gorm:"primaryKey"`
    Phone           string    `gorm:"unique;not null"`
    Name            string    `gorm:"not null"`
    Image           *string   `gorm:"null"` // Path to the uploaded image
    UserType        int       `gorm:"not null"`
    IsOtpVerified   bool      `gorm:"default:false"`
    GovIdType       *string   `gorm:"type:enum('aadhar', 'gst');default:null"`
    GovIdVal        *int      `gorm:"null"`
    IsGovIdVerified bool      `gorm:"default:false"`
    House           *string   `gorm:"null"`
    Area            *string   `gorm:"null"`
    Landmark        *string   `gorm:"null"`
    SdAddressID     int       `gorm:"default:0"`
    CreatedAt       time.Time `gorm:"autoCreateTime"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
