package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             string
	Email          string
	Password       string
	Fullname       string
	ProfilePicture string
	Birthdate      string
	Gender         string
	BloodType      string
	Height         int
	Weight         int
	Role           string
	OTP            string
	IsVerified     bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
