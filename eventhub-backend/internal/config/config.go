package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JwtSecretKey string
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("No .env file found")
	}
	return Config{
		JwtSecretKey: getEnv("JWT_SECRET_KEY", "mysecret"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
