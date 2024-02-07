package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             string `gorm:"primarykey"`
	Fullname       string
	Email          string `gorm:"not null"`
	Password       string `gorm:"not null"`
	ProfilePicture string
	Birthdate      string
	Gender         string `gorm:"type:enum('male', 'female');default:null"`
	BloodType      string `gorm:"type:enum('A', 'B', 'O', 'AB');default:null"`
	Height         int
	Weight         int
	Role           string `gorm:"type:enum('user');default:'user'"`
	OTP            string `gorm:"not null"`
	OTPExpiration  int64
	VerifyAccount  string
	IsVerified     bool `gorm:"not null;default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
