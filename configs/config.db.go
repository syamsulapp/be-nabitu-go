package configs

import (
	"fmt"
	"os"
)

var (
	DB_HOSTS    = os.Getenv("DB_HOSTS")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_DATABASE = os.Getenv("DB_DATABASE")
)

func InitConfigDb() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", DB_USERNAME, DB_PASSWORD, DB_HOSTS, DB_DATABASE)
}
