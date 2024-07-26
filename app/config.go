package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host        string
	Port        string
	DatabaseURL string
	Prefork     bool
}

func FromEnv() Config {
	dbUrl := os.Getenv("DATABASE_URL")
	prefork := os.Getenv("PREFORK")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	return Config{
		Host:        host,
		Port:        port,
		DatabaseURL: dbUrl,
		Prefork:     prefork == "1",
	}
}

func (c Config) URL() string {
	return c.Host + ":" + c.Port
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
