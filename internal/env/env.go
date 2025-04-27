package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	//"strconv"
)

func GetString(key string, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valInt

}
