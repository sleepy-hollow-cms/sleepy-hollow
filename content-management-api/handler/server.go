package handler

import (
	"content-management-api/cache"
	"content-management-api/driver/mongo"
	"content-management-api/env"
	"content-management-api/gateway"
	"content-management-api/usecase"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type instance struct {
	port   int
	server *echo.Echo
}

type Server interface {
	Start()
}

func (e *instance) Start() {
	e.server.Logger.Fatal(e.server.Start(fmt.Sprintf(":%v", e.port)))
}

func NewServer(container cache.Cache) Server {
	config := env.GetServerConfig()
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/v1/systems/ping", pong)

	routing(e, container)

	return &instance{
		port:   config.Port,
		server: e,
	}
}

func routing(e *echo.Echo, container cache.Cache) *echo.Echo {
	// TODO move to outside of handler
	db, err := container.Load(cache.MongoDB)
	if err != nil {
		return nil
	}
	mongoContentModelDriver := mongo.NewContentModelDriver(db.(*mongo.Client))
	mongoEntryDriver := mongo.NewEntryDriver(db.(*mongo.Client))

	contentModelResource := NewContentModelResource(usecase.NewContentModel(gateway.NewContentModel(mongoContentModelDriver)))
	entryResource := NewEntryResource(usecase.NewEntry(gateway.NewEntry(mongoEntryDriver)))
	spaceResource := NewSpaceResource(usecase.NewSpace(gateway.NewSpace()))

	contentModelResource.Routing(e)
	entryResource.Routing(e)
	spaceResource.Routing(e)

	return e
}
