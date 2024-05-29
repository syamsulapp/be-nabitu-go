package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func InitConfigEnv() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	envPath := filepath.Join(dir, "../.env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("err because didn't load .env")
	}
}
