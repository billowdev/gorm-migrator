package utils

import (
	"time"

	"github.com/billowdev/gorm-migrator/internal/models"
	"gorm.io/gorm"
)

// EnsureMigrationHistoryTable creates the migration history table if it doesn't exist.
func EnsureMigrationHistoryTable(db *gorm.DB) error {
	return models.CreateMigrationHistoryTable(db)
}

// IsMigrationApplied checks if a migration has already been applied.
func IsMigrationApplied(db *gorm.DB, migrationID string) (bool, error) {
	var count int64
	if err := db.Model(&models.MigrationHistory{}).Where("id = ?", migrationID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// LogMigration records the applied migration in the history.
func LogMigration(db *gorm.DB, migration models.Migration) error {
	history := models.MigrationHistory{
		ID:          migration.ID,
		AppliedAt:   time.Now(),
		Description: migration.Description,
	}
	return db.Create(&history).Error
}
