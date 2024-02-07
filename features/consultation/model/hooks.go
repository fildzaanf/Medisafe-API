package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (consultation *Consultation) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	consultation.ID = UUID.String()

	return nil
}
