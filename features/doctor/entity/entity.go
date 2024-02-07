package entity

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID               string
	Fullname         string
	Email            string
	Password         string
	ConfirmPassword  string
	NewPassword      string
	ProfilePicture   string
	Gender           string
	Status           bool
	Price            int
	Specialist       string
	Experience       string
	NoSTR            int
	Role             string
	Alumnus          string
	AboutDoctor      string
	LocationPractice string
	OTP              string
	OTPExpiration    int64
	VerifyAccount    string
	IsVerified       bool
	UpdatedAt        time.Time
	CreatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}
