package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/usecase"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/util"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/util/log"
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
	g.POST("/spaces/:spaceId/entries", en.CreateEntry)
	g.GET("/spaces/:spaceId/entries", en.FindEntries)
	g.GET("/spaces/:spaceId/entries/:entryId", en.FindEntry)
	g.DELETE("/spaces/:spaceId/entries/:entryId", en.DeleteEntry)

	g.PUT("/spaces/:spaceId/entries/:entryId/published", en.PublishedEntry)
	g.DELETE("/spaces/:spaceId/entries/:entryId/published", en.UnPublishedEntry)

	g.PUT("/spaces/:spaceId/entries/:entryId/archived", en.ArchivedEntry)
	g.DELETE("/spaces/:spaceId/entries/:entryId/archived", en.UnArchivedEntry)
}

func (en *EntryResource) FindEntries(c echo.Context) error {
	queryContentModelID := c.QueryParam("contentModelId")

	entries, err := en.EntryUseCase.Find(domain.ContentModelID(queryContentModelID))

	if err != nil {
		switch err.(type) {
		case usecase.EntryNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	response := make([]EntryPostResponseBody, len(entries))
	for i, entry := range entries {
		itemsResponse := make([]ItemsRequestBody, len(entry.Items))
		for j, item := range entry.Items {
			itemsResponse[j] = ItemsRequestBody{
				Value: item.Value,
			}
		}

		response[i] = EntryPostResponseBody{
			ID:             entry.ID.String(),
			ContentModelID: entry.ContentModelID.String(),
			Items:          itemsResponse,
		}
	}

	return c.JSON(http.StatusOK, response)
}

func (en *EntryResource) FindEntry(c echo.Context) error {
	entryId := c.Param("entryId")

	findEntry, err := en.EntryUseCase.FindByID(domain.EntryId(entryId))

	if err != nil {
		switch err.(type) {
		case usecase.EntryNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	items := make([]ItemsRequestBody, len(findEntry.Items))
	for i, item := range findEntry.Items {
		items[i] = ItemsRequestBody{
			Value: item.Value,
		}
	}

	responseBody := EntryPostResponseBody{
		ID:             findEntry.ID.String(),
		ContentModelID: findEntry.ContentModelID.String(),
		Items:          items,
	}

	return c.JSON(http.StatusOK, responseBody)
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
		value, err := domain.FactoryValue(item.Value)
		if err != nil {
			errStacks.Collect(err)
		}
		entryItems[i] = domain.EntryItem{
			Value: value,
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
			log.Logger.Warnf("Entry cannot register becouse Content Model ID %s not found", requestBody.ContentModelID)
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		case domain.EntryValidationError:
			log.Logger.Warnf("Entry cannot register becouse shape of Entry dose not match to Content Model.\nCause: %s", err.Error())
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
			Value: entryItem.Value,
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

func (en *EntryResource) DeleteEntry(c echo.Context) error {
	entryId := c.Param("entryId")

	err := en.EntryUseCase.DeleteByID(domain.EntryId(entryId))

	if err != nil {
		switch err.(type) {
		case usecase.EntryNotFoundError:
			return c.String(http.StatusNotFound, err.Error())
		default:
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	return c.NoContent(http.StatusNoContent)
}

func (en *EntryResource) PublishedEntry(c echo.Context) error {
	_ = c.Param("entryId")
	return c.JSON(http.StatusInternalServerError, nil)
}

func (en *EntryResource) UnPublishedEntry(c echo.Context) error {
	_ = c.Param("entryId")
	return c.JSON(http.StatusInternalServerError, nil)
}

func (en *EntryResource) ArchivedEntry(c echo.Context) error {
	_ = c.Param("entryId")
	return c.JSON(http.StatusInternalServerError, nil)
}

func (en *EntryResource) UnArchivedEntry(c echo.Context) error {
	_ = c.Param("entryId")
	return c.JSON(http.StatusInternalServerError, nil)
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
