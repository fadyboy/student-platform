package dbclient

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // import for sqlite3 driver
	"github.com/student-plaform/studentservice/models"
)

// InitializeDatabase - database initialization with the student table
func InitializeDatabase() {
	db := ConnectToDatabase()
	defer db.Close()

	// Migrate the database schema
	db.AutoMigrate(&models.Student{})
}

// ConnectToDatabase - Establishes a database connection
func ConnectToDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "student.db")
	if err != nil {
		log.Panic("Error connecting to the database")
	}
	return db

}
