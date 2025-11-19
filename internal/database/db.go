package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to PostgreSQL database")
	
	// Optional auto-migration - enable with AUTO_MIGRATE=true
	if os.Getenv("AUTO_MIGRATE") == "true" {
		log.Println("Running auto-migration...")
		AutoMigrate()
	} else {
		log.Println("Skipping auto-migration (set AUTO_MIGRATE=true to enable)")
	}
}
