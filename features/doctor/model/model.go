package model

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID               string `gorm:"primarykey"`
	Fullname         string `gorm:"not null"`
	Email            string `gorm:"not null"`
	Password         string `gorm:"not null"`
	ConfirmPassword  string `gorm:"not null"`
	ProfilePicture   string `gorm:"not null"`
	Gender           string `gorm:"type:enum('male', 'female')"`
	Status           bool   `gorm:"not null;default:false"`
	Price            int    `gorm:"type:enum('Rp100000');default:'Rp100000'"`
	Specialist       string `gorm:"not null"`
	Experience       string `gorm:"not null"`
	NoSTR            int    `gorm:"not null"`
	Role             string `gorm:"type:enum('doctor');default:'doctor'"`
	Alumnus          string `gorm:"not null"`
	AboutDoctor      string `gorm:"not null"`
	LocationPractice string `gorm:"not null"`
	OTP              string `gorm:"not null"`
	OTPExpiration    int64
	VerifyAccount    string
	IsVerified       bool `gorm:"not null;default:false"`
	UpdatedAt        time.Time
	CreatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
