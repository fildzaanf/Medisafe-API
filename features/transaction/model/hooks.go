package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (transaction *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	transaction.ID = UUID.String()

	return nil
}
