package handler

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/usecase"
	"content-management-api/usecase/write"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
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
	g.DELETE("/spaces/:spaceId/contentModels/:contentModelId", r.DeleteByID)
	g.GET("/spaces/:spaceId/contentModels", r.GetBySpaceID)
	g.POST("/spaces/:spaceId/contentModels", r.CreateContentModel)
}

func (r *ContentModelResource) GetByID(c echo.Context) error {
	contentModelId := c.Param("contentModelId")

	contentModel, err := r.ContentModelUseCase.FindByID(domain.ContentModelID(contentModelId))

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	resFields := make([]Field, len(contentModel.Fields))
	for i, field := range contentModel.Fields {
		resFields[i] = Field{
			Name:     field.Name.String(),
			Type:     field.Type.String(),
			Required: bool(field.Required),
		}
	}

	return c.JSON(http.StatusOK, ContentModelResponseBody{
		ID:     contentModel.ID.String(),
		Name:   contentModel.Name.String(),
		Fields: resFields,
	})
}

func (r *ContentModelResource) DeleteByID(c echo.Context) error {
	contentModelId := c.Param("contentModelId")

	err := r.ContentModelUseCase.DeleteContentModelByID(domain.ContentModelID(contentModelId))

	if err != nil {
		switch {
		case errors.As(err, &usecase.ContentModelNotFoundError{}):
			return c.String(http.StatusNotFound, err.Error())
		default:
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	return c.NoContent(http.StatusNoContent)
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
	m := ContentModelPostRequestBody{}

	if err := c.Bind(&m); err != nil {
		c.String(http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(m); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return err
	}

	fields := make(field.Fields, len(m.Fields))
	for i, f := range m.Fields {
		fields[i] = field.Field{
			Type:     field.Of(f.Type),
			Required: field.Required(f.Required),
			Name:     field.Name(f.Name),
		}
	}

	contentModel, err := r.ContentModelUseCase.Create(write.ContentModel{
		Name:   domain.Name(m.Name),
		Fields: fields,
	})

	if err != nil {
		println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	resFields := make([]Field, len(contentModel.Fields))
	for i, field := range contentModel.Fields {
		resFields[i] = Field{
			Name:     field.Name.String(),
			Type:     field.Type.String(),
			Required: bool(field.Required),
		}
	}

	c.JSON(http.StatusCreated, ContentModelResponseBody{
		ID:     contentModel.ID.String(),
		Name:   contentModel.Name.String(),
		Fields: resFields,
	})

	return nil
}

type ContentModelResponseBody struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type ContentModelPostRequestBody struct {
	Name   string  `json:"name" validate:"required"`
	Fields []Field `json:"fields" validate:"dive,required"`
}

type Field struct {
	Type     string `json:"type" validate:"required"`
	Required bool   `json:"required"`
	Name     string `json:"name"`
}
