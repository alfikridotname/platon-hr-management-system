package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupConnection is func for setup connection to database
func SetupConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error load file env")
	}

	// Init DB
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Connect DB
	dsn := fmt.Sprintf("host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Jakarta")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Check Error
	if err != nil {
		panic("Error connect to database")
	}

	return db
}

// CloseConnection is func for close connection to database
func CloseConnection(db *gorm.DB) {
	dbPG, err := db.DB()

	if err != nil {
		panic("Failed to close database connection")
	}

	dbPG.Close()
}
