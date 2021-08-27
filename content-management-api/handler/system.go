package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
