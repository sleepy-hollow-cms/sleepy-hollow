package handler

import (
	"content-management-api/domain"
	"content-management-api/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	ContentModel struct {
		ID string `json:"id"`
	}
)

type ContentModelResource struct {
	ContentModelUseCase *usecase.ContentModel
}

func NewContentModelResource(useCase *usecase.ContentModel) *ContentModelResource {
	return &ContentModelResource{
		ContentModelUseCase: useCase,
	}
}

func (r *ContentModelResource) Routing(e *echo.Echo) {
	g := e.Group("/v1")
	g.GET("/spaces/:spaceId/contentModels/:contentModelId", r.GetByID)
	g.GET("/spaces/:spaceId/contentModels", r.GetBySpaceID)
}

func (r *ContentModelResource) GetByID(c echo.Context) error {
	contentModelId := c.Param("contentModelId")

	contentModel, err := r.ContentModelUseCase.FindContentModelByID(domain.ContentModelID(contentModelId))

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &ContentModel{
		ID: contentModel.ID.String(),
	})
}

func (r *ContentModelResource) GetBySpaceID(c echo.Context) error {
	spaceId := c.Param("spaceId")

	contentModels, err := r.ContentModelUseCase.FindContentModelBySpaceID(domain.SpaceID(spaceId))

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	rest := make([]ContentModel, len(contentModels))
	for i, contentModel := range contentModels {
		rest[i] = ContentModel{
			ID: contentModel.ID.String(),
		}
	}

	return c.JSON(http.StatusOK, rest)
}
