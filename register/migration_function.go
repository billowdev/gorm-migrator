package register

import (
	"github.com/billowdev/gorm-migrator/internal/models"
	"gorm.io/gorm"
)

var MigrationID_M20240921T1552_0001_AddColumnName = "M20240921T1552_0001_AddColumnName"

func M20240921T1552_0001_AddColumnName() models.Migration {
	return models.Migration{
		ID:          MigrationID_M20240921T1552_0001_AddColumnName,
		Description: "Adds name column to users, backs up order_information_id, and sets foreign key constraint.",
		Up: func(db *gorm.DB) error {
			return db.Transaction(func(tx *gorm.DB) error {
				sql := `ALTER TABLE users ADD COLUMN name VARCHAR;`
				if err := tx.Exec(sql).Error; err != nil {
					return err
				}
				return nil
			})
		},
		Down: func(db *gorm.DB) error {
			return db.Transaction(func(tx *gorm.DB) error {
				sql := `ALTER TABLE users DROP COLUMN IF EXISTS name;`
				if err := tx.Exec(sql).Error; err != nil {
					return err
				}
				return nil
			})
		},
	}
}
