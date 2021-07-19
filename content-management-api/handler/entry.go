package handler

import (
	"content-management-api/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EntryResource struct {
	EntryUseCase *usecase.Entry
}

func NewEntryResource(useCase *usecase.Entry) *EntryResource {
	return &EntryResource{
		EntryUseCase: useCase,
	}
}

func (en *EntryResource) Routing(e *echo.Echo) {
	g := e.Group("/v1")
	g.POST("/specs/:spaceId/contentModels/:modelId/entry", en.CreateEntry)
}

func (en *EntryResource) CreateEntry(c echo.Context) error {
	entry, err := en.EntryUseCase.Create()

	if err != nil {
		println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusCreated, EntryPostResponseBody{
		ID: entry.ID.String(),
	})

	return nil
}

type EntryPostResponseBody struct {
	ID string `json:"id"`
}
