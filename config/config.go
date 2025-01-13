package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App App
	Db  Db
}

type App struct {
	Port string
}

type Db struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func New() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		App: App{
			Port: os.Getenv("APP_PORT"),
		},
		Db: Db{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}
}
