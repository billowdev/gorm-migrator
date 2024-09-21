package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/billowdev/gorm-migrator/internal/core"
	"github.com/billowdev/gorm-migrator/register"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Define flags for rollback, migrate up, and specific migration file
	rollback := flag.String("rollback", "", "Specify the migration ID to roll back")
	migrateUp := flag.Bool("migrate-up", false, "Run migrations (migrate up)")
	migrateFile := flag.String("up", "", "Specify a specific migration file to apply")
	flag.Parse()

	// Retrieve DSN from environment variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s search_path=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
		os.Getenv("DB_SCHEMA"),
	)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Check if a rollback is requested
	if *rollback != "" {
		if err := core.Rollback(db, *rollback, register.AllMigrationFunctions); err != nil {
			log.Fatalf("rollback failed: %v", err)
		}
		log.Printf("Rollback of migration %s successful!", *rollback)
		return // Exit after rollback
	}

	// Run migrations if the flag is set
	if *migrateUp {
		if *migrateFile != "" {
			if err := core.MigrateSpecific(db, *migrateFile, register.AllMigrationFunctions); err != nil {
				log.Fatalf("migration of %s failed: %v", *migrateFile, err)
			}
			log.Printf("Migration of %s applied successfully!", *migrateFile)
		} else {
			if err := core.Migrate(db, register.AllMigrationFunctions); err != nil {
				log.Fatalf("migration failed: %v", err)
			}
			log.Println("Migrations applied successfully!")
		}
		return // Exit after migration
	}

	// If no flags are set, you can provide default behavior or help
	log.Println("No flags set. Use -migrate-up to apply migrations or -rollback to undo a migration.")
}
