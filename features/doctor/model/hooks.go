package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (doctor *Doctor) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	doctor.ID = UUID.String()

	return nil
}
