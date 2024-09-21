package register

import "github.com/billowdev/gorm-migrator/internal/models"

var phase1Migrations = []models.Migration{
	M20240921T1552_0001_AddColumnName(),
}

var nextPhaseMigrations = []models.Migration{}

var AllMigrationFunctions = append(phase1Migrations, nextPhaseMigrations...)
