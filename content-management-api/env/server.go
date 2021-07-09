package env

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type ServerConfig struct {
	Port int
}

func GetServerConfig() *ServerConfig {
	db := ServerConfig{}
	err := envconfig.Process("app", &db)

	if err != nil {
		log.Fatal(err)
	}

	return &db
}
