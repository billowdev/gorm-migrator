package configs

import (
	"log"
	"os"

	"github.com/billowdev/gorm-migrator/pkg/utils"
	"github.com/joho/godotenv"
)

var (
	DB_DRY_RUN       bool
	DB_NAME          string
	DB_USERNAME      string
	DB_PASSWORD      string
	DB_HOST          string
	DB_PORT          string
	DB_SSL_MODE      string
	DB_RUN_MIGRATION bool
	DB_RUN_SEEDER    bool
	DB_SCHEMA        string
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}
	DB_DRY_RUN = false

	DB_NAME = os.Getenv("DB_NAME")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_SSL_MODE = os.Getenv("DB_SSL_MODE")
	if DB_SSL_MODE == "" {
		DB_SSL_MODE = "default"
	}
	DB_SCHEMA = os.Getenv("DB_SCHEMA")

	DB_RUN_MIGRATION = utils.ParseBoolEnv("DB_RUN_MIGRATION", false)
	DB_RUN_SEEDER = utils.ParseBoolEnv("DB_RUN_SEEDER", false)
}
