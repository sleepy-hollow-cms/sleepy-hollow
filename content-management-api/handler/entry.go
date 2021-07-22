package handler

import (
	"content-management-api/domain"
	"content-management-api/usecase"
	"content-management-api/usecase/write"
	"content-management-api/util/log"
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
	g.POST("/specs/:spaceId/entries", en.CreateEntry)
}

func (en *EntryResource) CreateEntry(c echo.Context) error {
	modelID := EntryPostRequestBody{}

	if err := c.Bind(&modelID); err != nil {
		c.String(http.StatusBadRequest, "invalid request body")
		return err
	}

	entry := write.Entry{
		ContentModelID: domain.ContentModelID(modelID.ContentModelID),
	}

	createdEntry, err := en.EntryUseCase.Create(entry)

	if err != nil {
		switch err := err.(type) {
		case usecase.ContentModelNotFoundError:
			log.Logger.Warnf("Entry cannot Found Becouse Content Model ID %s Not Found", modelID.ContentModelID)
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		default:
			log.Logger.Warnf("Something Happened: %s", modelID.ContentModelID)
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		return nil
	}

	c.JSON(http.StatusCreated, EntryPostResponseBody{
		ID:             createdEntry.ID.String(),
		ContentModelID: createdEntry.ContentModelID.String(),
	})

	return nil
}

type EntryPostResponseBody struct {
	ID             string `json:"id"`
	ContentModelID string `json:"content-model-id"`
}

type EntryPostRequestBody struct {
	ContentModelID string `json:"content-model-id"`
}