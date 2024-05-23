package configs

import (
	"fmt"
)

func InitConfigDb(username string, password string, hosts string, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hosts, database)
}
