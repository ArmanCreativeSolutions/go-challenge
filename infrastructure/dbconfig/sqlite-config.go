package dbconfig

import (
	_ "github.com/mattn/go-sqlite3"
	"go-challenge/domain/usersegments"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func Initialize(dbfileName string) {
	err := os.Remove(dbfileName)
	if err != nil {
		log.Println("error removing file", err)
	} else {
		log.Println("successfully deleted file")
	}

	db, err := gorm.Open(sqlite.Open(dbfileName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("database connected...")

	migrationErr := db.AutoMigrate(&usersegments.UserSegment{})
	if migrationErr != nil {
		log.Fatal(migrationErr)
	}
	log.Println("migrated successfully")
}
