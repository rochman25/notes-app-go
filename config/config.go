package config

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Env not found")
	}
	return os.Getenv(key)
}
