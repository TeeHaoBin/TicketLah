package db

import (
	"gorm.io/gorm"
	"github.com/TeeHaoBin/TicketLah/backend/models"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Event{},
	)
}