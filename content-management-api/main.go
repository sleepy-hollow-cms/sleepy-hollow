package main

import (
	"content-management-api/handler"
)

func main() {
	server := handler.NewServer()
	server.Start()
}
