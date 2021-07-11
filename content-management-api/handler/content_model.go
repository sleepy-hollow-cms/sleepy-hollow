package handler

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/usecase"
	"content-management-api/usecase/write"
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
	g.PUT("/spaces/:spaceId/contentModels", r.CreateContentModel)
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

func (r *ContentModelResource) CreateContentModel(c echo.Context) error {
	m := ContentModelPutRequestBody{}

	if err := c.Bind(&m); err != nil {
		c.String(http.StatusBadRequest, "invalid request body")
		return err
	}

	fields := make(field.Fields, len(m.Fields))
	for i, f := range m.Fields {
		fields[i] = field.Field{Type: field.Of(f.Type)}
	}

	contentModel, err := r.ContentModelUseCase.Create(write.ContentModel{
		Fields: fields,
	})

	if err != nil {
		println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	resFields := make([]Field, len(contentModel.Fields))
	for i, field := range contentModel.Fields {
		resFields[i] = Field{
			Type: field.Type.String(),
		}
	}

	c.JSON(http.StatusOK, ContentModelPutResponseBody{
		ID:     contentModel.ID.String(),
		Name:   "NOT IMPLEMENTED",
		Fields: resFields,
	})

	return nil
}

type ContentModelPutResponseBody struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type ContentModelPutRequestBody struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Type string `json:"type"`
}