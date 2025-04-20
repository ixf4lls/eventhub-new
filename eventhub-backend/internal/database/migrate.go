package database

import (
	"eventhub-backend/internal/repository"

	"gorm.io/gorm"
)

func MigrageDB(DB *gorm.DB) {
	DB.AutoMigrate(
		&repository.EventModel{},
		&repository.UserModel{},
		&repository.OrganizationModel{},
	)
}
