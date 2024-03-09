package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectMemoryDb() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	return

}

func ClosedConnectionMemoryDb(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.Close()

	if err != nil {
		log.Fatal(err)
	}
}
