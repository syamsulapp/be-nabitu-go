package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitConnectionDBMysql(dsn string) {
	var err error

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = d

	log.Println("DB Successfully Connect")
}

func GetDBMysql() *gorm.DB {
	return db
}
