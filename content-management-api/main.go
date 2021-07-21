package main

import (
	"content-management-api/cache"
	"content-management-api/driver/mongo"
	"content-management-api/handler"
	"content-management-api/util/log"
)

func main() {
	log.Logger.Debugw("Starting...")

	container := cache.NewContainer()
	dbClient := mongo.NewClient()
	_, err := dbClient.Connect()
	if err != nil {
		log.Logger.Fatalw("DB Connect Error", log.Logger.Error(err))
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
