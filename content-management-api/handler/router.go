package handler

import (
	"content-management-api/gateway"
	"content-management-api/usecase"
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
	routing(e)

	return &instance{
		server: e,
	}
}

func routing(e *echo.Echo) *echo.Echo {
	contentModelResource := NewContentModelResource(
		usecase.NewContentModel(gateway.NewContentModel()),
	)

	v1Routing(e, contentModelResource)

	return e
}

func v1Routing(e *echo.Echo, resource *ContentModelResource) *echo.Echo {
	g := e.Group("/v1")
	g.GET("/spaces/:spaceId/contentModels/:contentModelId", resource.GetByID)
	g.GET("/spaces/:spaceId/contentModels", resource.GetBySpaceID)
	return e
}
