package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              string
	Email           string
	Password        string
	NewPassword     string
	ConfirmPassword string
	Fullname        string
	ProfilePicture  string
	Birthdate       string
	Gender          string
	BloodType       string
	Height          int
	Weight          int
	Role            string
	OTP             string
	OTPExpiration   int64
	VerifyAccount   string
	IsVerified      bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
