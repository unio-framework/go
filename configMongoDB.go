package unio

import (
	"github.com/labstack/gommon/log"
    "github.com/zebresel-com/mongodm"
)

/**
Include additional information, like register tables
 */
type MongoConfig func(connection *mongodm.Connection)

func (c *Config) MongoConnection(afterConfig MongoConfig) *mongodm.Connection {
	connection, err := mongodm.Connect(config()); if err != nil {
		log.Fatal(err)
	}
	if afterConfig != nil { afterConfig(connection) }
	return connection
}

func config() *mongodm.Config {
    config := &mongodm.Config{
        DatabaseHosts: []string{ Configs.Env("MONGODB_HOST") },
        DatabaseName: Configs.Env("DATABASE_NAME"),
    }
    if Configs.Env("MONGODB_USERNAME") != "" { config.DatabaseUser = Configs.Env("MONGODB_USERNAME") }
    if Configs.Env("MONGODB_PASSWORD") != "" { config.DatabasePassword = Configs.Env("MONGODB_PASSWORD") }
    if Configs.Env("MONGODB_SOURCE") != "" { config.DatabaseSource = Configs.Env("MONGODB_SOURCE") }

    return config
}