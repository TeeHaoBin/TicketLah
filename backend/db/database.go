package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"
	"github.com/TeeHaoBin/TicketLah/backend/config"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func Init(cfg *config.EnvConfig, DBMigrator func(*gorm.DB) error) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Info("Database connected successfully")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Unable to migrate database: %v", err)
	}

	log.Info("Database migrated successfully")

	return db
}
