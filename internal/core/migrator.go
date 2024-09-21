package core

import (
	"errors"
	"fmt"

	"github.com/billowdev/gorm-migrator/internal/models"
	"github.com/billowdev/gorm-migrator/pkg/utils"
	"gorm.io/gorm"
)

// Migrate applies all migrations that have not been applied yet.
func Migrate(db *gorm.DB, migrationFunctions []models.Migration) error {
	if err := utils.EnsureMigrationHistoryTable(db); err != nil {
		return err
	}

	for _, migration := range migrationFunctions {
		applied, err := utils.IsMigrationApplied(db, migration.ID)
		if err != nil {
			return err
		}
		if !applied {
			if err := migration.Up(db); err != nil {
				return err
			}
			// Log the applied migration
			if err := utils.LogMigration(db, migration); err != nil {
				return err
			}
		}
	}
	return nil
}

// MigrateSpecific applies a specific migration
func MigrateSpecific(db *gorm.DB, migrationID string, migrationFunctions []models.Migration) error {
	if err := utils.EnsureMigrationHistoryTable(db); err != nil {
		return err
	}

	var targetMigration *models.Migration
	for _, migration := range migrationFunctions {
		if migration.ID == migrationID {
			targetMigration = &migration
			break
		}
	}

	if targetMigration == nil {
		return errors.New("migration not found")
	}

	applied, err := utils.IsMigrationApplied(db, targetMigration.ID)
	if err != nil {
		return err
	}

	if applied {
		return errors.New("migration has already been applied")
	}

	if err := targetMigration.Up(db); err != nil {
		return err
	}

	if err := utils.LogMigration(db, *targetMigration); err != nil {
		return err
	}

	return nil
}

// Rollback undoes a specific migration.
func Rollback(db *gorm.DB, migrationID string, migrationFunctions []models.Migration) error {
	var history models.MigrationHistory
	if err := db.First(&history, "id = ?", migrationID).Error; err != nil {
		return fmt.Errorf("migration not found: %s", migrationID)
	}

	for _, migration := range migrationFunctions {
		if migration.ID == migrationID {
			if err := migration.Down(db); err != nil {
				return err
			}
			// Remove the migration from history
			return db.Delete(&history).Error
		}
	}
	return fmt.Errorf("migration not found: %s", migrationID)
}
