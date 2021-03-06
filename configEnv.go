package unio

import (
	"github.com/joho/godotenv"
    "github.com/labstack/gommon/log"
    "os"
)

/**
Load enviroment variables
*/
func (c *Config) LoadEnv() {
    baseFiles := []string{
        ".env",
        ".env."+c.Environment(),
    }
    for _,file := range baseFiles {
        err := godotenv.Overload(file); if err != nil {
            log.Error("Error loading "+file+" file")
        }
    }
}

/**
Shortcut to get an environment
*/
func (c *Config) Env(key string) string {
	return os.Getenv(key)
}

/**
For project environment, need to set GOENV with environment type
 */
func (c *Config) Environment() string {
    env := os.Getenv("GOENV")
    if env == "" { env = "development" }
    return env
}