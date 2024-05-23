package configs

import (
	"fmt"
)

func InitConfigDbMysql(username string, password string, hosts string, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", username, password, hosts, database)
}
