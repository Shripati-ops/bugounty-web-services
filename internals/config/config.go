package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName string
	HTTP        HTTPConfig
	DB          DBConfig
}

type HTTPConfig struct {
	Port int64
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
}

func LoadConfig() (*Config, error) {
	initEnv()
	return &Config{
		ServiceName: "bugounty",
		HTTP: HTTPConfig{
			Port: 8080,
		},
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
	}, nil
}
