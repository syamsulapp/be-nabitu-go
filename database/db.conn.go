package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitConnectionDB(dsn string) {
	var err error

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = d

	log.Println("Db successfully connect")
}

func GetDB() *gorm.DB {
	return db
}
