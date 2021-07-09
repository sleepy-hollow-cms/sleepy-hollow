package main

import (
	"content-management-api/cache"
	"content-management-api/driver/mongo"
	"content-management-api/handler"
	"log"
)

func main() {
	container := cache.NewContainer()
	db, err := mongo.NewClient().Connect()

	if err != nil {
		log.Fatal(err)
	}

	container.Store(cache.MongoDB, db)
	server := handler.NewServer(container)
	server.Start()
}
