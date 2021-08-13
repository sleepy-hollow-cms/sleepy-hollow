package handler

import (
	"content-management-api/domain"
	"content-management-api/usecase"
	"content-management-api/util"
	"content-management-api/util/log"
	"fmt"
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

	entryItems := make([]domain.EntryItem, len(requestBody.Items))
	errStacks := new(util.ErrorCollector)
	for i, item := range requestBody.Items {
		contentType := domain.Of(item.ContentType)
		value, err := domain.FactoryValue(contentType, item.Value)
		if err != nil {
			errStacks.Collect(err)
		}
		entryItems[i] = domain.EntryItem{
			Type:      contentType,
			FieldName: domain.Name(item.Name),
			Value:     value,
		}
	}

	if errStacks.Size() != 0 {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid request bod\n%s", errStacks.Error()))
		return errStacks
	}

	createdEntry, err := en.EntryUseCase.Register(domain.Entry{
		ContentModelID: domain.ContentModelID(requestBody.ContentModelID),
		Items:          entryItems,
	})

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

	items := make([]ItemsRequestBody, len(createdEntry.Items))
	for i, entryItem := range createdEntry.Items {
		items[i] = ItemsRequestBody{
			ContentType: entryItem.Type.String(),
			Name:        entryItem.FieldName.String(),
			Value:       entryItem.Value,
		}
	}

	responseBody := EntryPostResponseBody{
		ID:             createdEntry.ID.String(),
		ContentModelID: requestBody.ContentModelID,
		Items:          items,
	}

	c.JSON(http.StatusCreated, responseBody)

	return nil
}

type EntryPostResponseBody struct {
	ID             string             `json:"id"`
	ContentModelID string             `json:"content-model-id"`
	Items          []ItemsRequestBody `json:"items"`
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
