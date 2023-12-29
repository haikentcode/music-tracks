package db

import (
	"musictracks/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// InitDB initializes the database connection
func InitDB() {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the database to create the 'todos' table
	db.AutoMigrate(&models.Track{})
	db.AutoMigrate(&models.Artist{})
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}
