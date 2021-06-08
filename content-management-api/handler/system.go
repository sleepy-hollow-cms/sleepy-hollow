package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	Pong struct {
		Ping string `json:"ping"`
	}
)

func pong(c echo.Context) error {
	return c.JSON(http.StatusOK, &Pong{
		Ping: "pong",
	})
}
