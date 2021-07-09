package env

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type MongoDBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
}

func GetMongoConfig() *MongoDBConfig {
	db := MongoDBConfig{}
	err := envconfig.Process("mongo", &db)

	if err != nil {
		log.Fatal(err)
	}

	return &db
}
