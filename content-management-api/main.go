package main

import (
	"content-management-api/cache"
	"content-management-api/driver/mongo"
	"content-management-api/handler"
	"log"
)

func main() {
	container := cache.NewContainer()
	dbClient := mongo.NewClient()
	_, err := dbClient.Connect()
	if err != nil {
		log.Fatal(err)
	}

	go dbClient.StartWatch()
	defer dbClient.StopWatch()
	defer func() {
		closeErr := dbClient.Disconnect()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}()

	container.Store(cache.MongoDB, dbClient)
	server := handler.NewServer(container)
	server.Start()
}
