package migration

import (
	consultation "talkspace/features/consultation/model"
	doctor "talkspace/features/doctor/model"
	transaction "talkspace/features/transaction/model"
	user "talkspace/features/user/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&doctor.Doctor{},
		&transaction.Transaction{},
		&consultation.Consultation{},
	)
}
