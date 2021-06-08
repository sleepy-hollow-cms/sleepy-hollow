package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type instance struct {
	server *echo.Echo
}

type Server interface {
	Start()
}

func (e *instance) Start() {
	e.server.Logger.Fatal(e.server.Start(":3000"))
}

func NewServer() Server {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/v1/systems/ping", pong)

	return &instance{
		server: e,
	}
}
