package unio

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

/**
Load enviroment variables
*/
func (c *Config) LoadEnv() {
    envTyped := ".env."+c.Environment()
	err := godotenv.Load(".env", envTyped); if err != nil {
		log.Fatal("Error loading .env file")
	}
}

/**
Shortcut to get an environment
*/
func (c *Config) Env(key string) string {
	return os.Getenv(key)
}

/**
For project environment
 */
func (c *Config) Environment() string {
    env := os.Getenv("GOENV")
    if env == "" { env = "development" }
    return env
}