package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseUrl string
	JwtSecret   string
}

func getPort() string {
	port := os.Getenv("HTTP_PORT")
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("HTTP_PORT is not an int: %v\n", err)
	}
	return port
}

func getDatabaseUrl() string {
	dbUrl := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		5432,
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)
	return dbUrl
}

func getJwtSecret() string {
	return os.Getenv("JWT_SECRET")
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
	return &Config{
		Port:        getPort(),
		DatabaseUrl: getDatabaseUrl(),
		JwtSecret:   getJwtSecret(),
	}
}
