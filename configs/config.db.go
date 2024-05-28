package configs

import (
	"fmt"
	"os"
)

func InitConfigDbMysql() string {
	//declare local variable for init connection db from .env
	var (
		DB_HOSTS    = os.Getenv("DB_HOSTS")
		DB_USERNAME = os.Getenv("DB_USERNAME")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_DATABASE = os.Getenv("DB_DATABASE")
	)

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", DB_USERNAME, DB_PASSWORD, DB_HOSTS, DB_DATABASE)
}
