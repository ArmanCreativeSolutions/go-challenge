package dbconfig

import (
	_ "github.com/mattn/go-sqlite3"
	"go-challenge/domain/usersegments"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

const DB_NAME = "test.db"

func InitDB() {
	db := ConnectDB()
	err := db.AutoMigrate(&usersegments.UserSegment{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("migrated successfully")
}

func ConnectDB() *gorm.DB {
	err := os.Remove(DB_NAME)
	if err != nil {
		log.Println("error removing file", err)
	} else {
		log.Println("successfully deleted file")
	}

	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("database connected...")
	return db
}
