package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBDSN  string
	Env    string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system env vars")
	}

	return &Config{
		Port:  getEnv("PORT", "8080"),
		Env:   getEnv("ENV", "development"),
		DBDSN: getDBDSN(),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getDBDSN() string {
	return "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=" + os.Getenv("DB_SSLMODE")
}