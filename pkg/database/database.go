package database

import (
	"bank-app/pkg/helpers"

	"github.com/jinzhu/gorm"
)

// Create global variable
var DB *gorm.DB

// Create InitDatabase function
func InitDatabase() {
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=bank-app password=19660827 sslmode=disable")
	helpers.HandleErr(err)
	// Set up connection pool
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}
