package entity

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID                string
	Fullname          string
	Email             string
	Password          string
	ProfilePicture    string
	Gender            string
	Status            bool
	Price             int
	Specialist        string
	Experience        string
	NoSTR             int
	Role              string
	Alumnus           string
	AboutDoctor       string
	PracticeLocation  string
	OTP               string
	OTPExpiration     int64
	VerificationToken string
	IsVerified        bool
	UpdatedAt         time.Time
	CreatedAt         time.Time
	DeletedAt         gorm.DeletedAt
}
