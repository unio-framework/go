package unio

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

/**
Load enviroment variables
*/
func (c Configs) LoadEnv() {
	err := godotenv.Load(); if err != nil {
		log.Fatal("Error loading .env file")
	}
}

/**
Shortcut to get an environment
*/
func (c Configs) Env(key string) string {
	return os.Getenv(key)
}
