package models

import (
	"time"

	"gorm.io/gorm"
)

type Migration struct {
	ID          string
	Description string
	Up          func(*gorm.DB) error
	Down        func(*gorm.DB) error
}

type MigrationHistory struct {
	ID          string    `gorm:"primaryKey"` // Unique identifier for the migration
	AppliedAt   time.Time // Timestamp when the migration was applied
	Description string    // Brief description of what the migration does
}

func CreateMigrationHistoryTable(db *gorm.DB) error {
	return db.AutoMigrate(&MigrationHistory{})
}
