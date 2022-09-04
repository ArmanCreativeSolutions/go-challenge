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

type SqliteDBService struct {
	Db *gorm.DB
}

func NewSqliteDBService() *SqliteDBService {
	return &SqliteDBService{}
}

func (s *SqliteDBService) InitDB() (*gorm.DB, error) {
	//Instance := ConnectDB()
	err := os.Remove(DB_NAME)
	if err != nil {
		log.Println("error removing file", err)
	} else {
		log.Println("successfully deleted file")
	}

	s.Db, err = gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("database connected...")
	err = s.Db.AutoMigrate(&usersegments.UserSegment{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("migrated successfully")
	return s.Db, err
}

//
//func ConnectDB() *gorm.DB {
//
//	return db
//}

func (s *SqliteDBService) GetDB() *gorm.DB {
	return s.Db
}
