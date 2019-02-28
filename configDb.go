package unio

import (
	"github.com/go-bongo/bongo"
	"github.com/labstack/gommon/log"
)

func (c Config) Connection() *bongo.Connection {
	connection, err := bongo.Connect(config()); if err != nil {
		log.Fatal(err)
	}
	return connection
}

func config() *bongo.Config {
	return &bongo.Config{
		ConnectionString: Configs.Env("DATABASE_CONNECTION"),
		Database:         Configs.Env("DATABASE_NAME"),
	}
}
