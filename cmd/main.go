package main

import (
	"fmt"
	"log"

	"github.com/billowdev/gorm-migrator/internal/core"
	"github.com/billowdev/gorm-migrator/internal/models"
	"github.com/billowdev/gorm-migrator/pkg/configs"
	"github.com/billowdev/gorm-migrator/register"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database configuration
var db *gorm.DB

func main() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v search_path=%v",
		configs.DB_HOST,
		configs.DB_USERNAME,
		configs.DB_PASSWORD,
		configs.DB_NAME,
		configs.DB_PORT,
		configs.DB_SSL_MODE,
		configs.DB_SCHEMA,
	)
	// //Connect to the database
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// dsn := configs.DB_NAME // SQLite database file
	// db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("failed to connect to database: %v", err)
	// }

	// Perform automatic migration for the User model
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("auto migration failed: %v", err)
	}

	// Create a new Fiber app
	app := fiber.New()

	// Define endpoint to trigger migration
	app.Post("/migrate", migrateUpHandler)

	// Define endpoint to roll back migration
	app.Post("/migrate/rollback", migrateDownHandler)

	// Start the server
	log.Println("Server is running on :8081")
	log.Fatal(app.Listen(":8081"))

}

// migrateHandler triggers manual migrations
func migrateUpHandler(c *fiber.Ctx) error {
	if err := core.MigrateSpecific(db, register.MigrationID_M20240921T1552_0001_AddColumnName, register.AllMigrationFunctions); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Migration failed: " + err.Error())
	}
	return c.SendString("Migration successful")
}

// migrateHandler triggers manual migrations
func migrateDownHandler(c *fiber.Ctx) error {
	if err := core.Rollback(db, register.MigrationID_M20240921T1552_0001_AddColumnName, register.AllMigrationFunctions); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Migration failed: " + err.Error())
	}
	return c.SendString("Migration successful")
}
