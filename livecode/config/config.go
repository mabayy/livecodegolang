package config

import (
	"log"

	"github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = "5432"
	dbname   = "databaru"
	db       *gorm.DB
	err      error
)

func StartDB() {
	db, err := gorm.Open(postgres.Open(config), &gorm.config{})
	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}
	db.Debug().AutoMigrate(structs.gorm{})
}
