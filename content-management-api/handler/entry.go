package handler

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
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
	requestBody := EntryPostRequestBody{}

	if err := c.Bind(&requestBody); err != nil {
		c.String(http.StatusBadRequest, "invalid request body")
		return err
	}

	entryItems := make([]write.EntryItem, len(requestBody.Items))
	for i, item := range requestBody.Items {
		contentType := field.Of(item.ContentType)
		entryItems[i] = write.EntryItem{
			FieldName: field.Name(item.Name),
			Value:     field.FactoryValue(contentType, item.Value),
		}
	}

	entry := write.Entry{
		ContentModelID: domain.ContentModelID(requestBody.ContentModelID),
	}

	createdEntry, err := en.EntryUseCase.Register(entry, entryItems)

	if err != nil {
		switch err := err.(type) {
		case usecase.ContentModelNotFoundError:
			log.Logger.Warnf("Entry cannot Register Becouse Content Model ID %s Not Found", requestBody.ContentModelID)
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		default:
			log.Logger.Warnf("Something Happened: %s", requestBody.ContentModelID)
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
	ContentModelID string             `json:"content-model-id"`
	Items          []ItemsRequestBody `json:"items"`
}

type ItemsRequestBody struct {
	ContentType string      `json:"contentType"`
	Name        string      `json:"name"`
	Value       interface{} `json:"value"`
}
