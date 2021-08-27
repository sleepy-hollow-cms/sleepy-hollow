package main

import (
	"time"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/cache"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/mongo"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/handler"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/util/log"
)

func main() {

	time.Local = time.UTC

	log.Logger.Debugw("Starting...")

	container := cache.NewContainer()
	dbClient := mongo.NewClient()
	_, err := dbClient.Connect()
	if err != nil {
		log.Logger.Fatalw("DB client error", log.Logger.Error(err))
	}
	if err := dbClient.Ping(5 * time.Second); err != nil {
		log.Logger.Fatalw("DB Connecting PING Error", log.Logger.Error(err))
		panic(err)
	}

	go dbClient.StartWatch()
	defer dbClient.StopWatch()
	defer func() {
		closeErr := dbClient.Disconnect()
		if closeErr != nil {
			log.Logger.Fatalw("DB Close Error", log.Logger.Error(err))
		}
	}()

	container.Store(cache.MongoDB, dbClient)
	server := handler.NewServer(container)
	server.Start()
}
