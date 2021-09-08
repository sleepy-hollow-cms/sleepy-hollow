package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/usecase"
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
	g.POST("/spaces/:spaceId/contentModels", r.Create)
	g.PUT("/spaces/:spaceId/contentModels/:contentModelId", r.Update)
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
		ID:        contentModel.ID.String(),
		Name:      contentModel.Name.String(),
		CreatedAt: contentModel.CreatedAt.Time().Format(time.RFC3339),
		UpdatedAt: contentModel.UpdatedAt.Time().Format(time.RFC3339),
		Fields:    resFields,
	})
}

func (r *ContentModelResource) DeleteByID(c echo.Context) error {
	contentModelId := c.Param("contentModelId")

	err := r.ContentModelUseCase.DeleteContentModelByID(domain.ContentModelID(contentModelId))

	if err != nil {
		switch {
		case errors.As(err, &usecase.ContentModelNotFoundError{}):
			return c.String(http.StatusNotFound, err.Error())
		case errors.As(err, &usecase.ReferenceByEntryError{}):
			return c.String(http.StatusUnprocessableEntity, err.Error())
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
		return c.String(http.StatusInternalServerError, "Somethings happened")
	}

	rest := make([]ContentModelResponseBody, len(contentModels.Models))
	for i, m := range contentModels.Models {

		resFields := make([]Field, len(m.Fields))
		for i, field := range m.Fields {
			resFields[i] = Field{
				Name:     field.Name.String(),
				Type:     field.Type.String(),
				Required: bool(field.Required),
			}
		}
		rest[i] = ContentModelResponseBody{
			ID:        m.ID.String(),
			Name:      m.Name.String(),
			CreatedAt: m.CreatedAt.Time().Format(time.RFC3339),
			UpdatedAt: m.UpdatedAt.Time().Format(time.RFC3339),
			Fields:    resFields,
		}
	}
	return c.JSON(http.StatusOK, ContentModelsResponseBody{
		SpaceID: spaceId,
		Models:  rest,
	})
}

func (r *ContentModelResource) Create(c echo.Context) error {
	m := ContentModelCreateRequestBody{}

	if err := c.Bind(&m); err != nil {
		c.String(http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(m); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return err
	}

	fields := make(domain.Fields, len(m.Fields))
	for i, f := range m.Fields {
		fields[i] = domain.Field{
			Type:     domain.Of(f.Type),
			Required: domain.Required(f.Required),
			Name:     domain.Name(f.Name),
		}
	}

	contentModel, err := r.ContentModelUseCase.Create(domain.ContentModel{
		Name:      domain.Name(m.Name),
		Fields:    fields,
		CreatedAt: domain.CreatedAt(time.Now()),
		UpdatedAt: domain.UpdatedAt(time.Now()),
	})

	if err != nil {
		println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
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
		ID:        contentModel.ID.String(),
		Name:      contentModel.Name.String(),
		Fields:    resFields,
		CreatedAt: contentModel.CreatedAt.Time().Format(time.RFC3339),
		UpdatedAt: contentModel.UpdatedAt.Time().Format(time.RFC3339),
	})

	return nil
}

type ContentModelsResponseBody struct {
	SpaceID string                     `json:"id"`
	Models  []ContentModelResponseBody `json:"models"`
}
type ContentModelResponseBody struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Fields    []Field `json:"fields"`
	CreatedAt string  `json:"created-at"`
	UpdatedAt string  `json:"updated-at"`
}

type ContentModelCreateRequestBody struct {
	Name   string  `json:"name" validate:"required"`
	Fields []Field `json:"fields" validate:"dive,required"`
}

type ContentModelUpdateRequestBody struct {
	Name      string  `json:"name" validate:"required"`
	Fields    []Field `json:"fields" validate:"dive,required"`
	UpdatedAt string  `json:"updated-at" validate:"required"`
}

type Field struct {
	Type     string `json:"type" validate:"required"`
	Required bool   `json:"required"`
	Name     string `json:"name"`
}

func (r *ContentModelResource) Update(c echo.Context) error {

	contentModelId := c.Param("contentModelId")

	m := ContentModelUpdateRequestBody{}

	if err := c.Bind(&m); err != nil {
		c.String(http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(m); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return err
	}

	t, err := time.Parse(time.RFC3339, m.UpdatedAt)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return err
	}

	fields := make(domain.Fields, len(m.Fields))
	for i, f := range m.Fields {
		fields[i] = domain.Field{
			Type:     domain.Of(f.Type),
			Required: domain.Required(f.Required),
			Name:     domain.Name(f.Name),
		}
	}

	contentModel, err := r.ContentModelUseCase.Update(domain.ContentModel{
		ID:        domain.ContentModelID(contentModelId),
		Name:      domain.Name(m.Name),
		Fields:    fields,
		UpdatedAt: domain.UpdatedAt(t),
	})

	if err != nil {
		println(err)
		switch {
		case errors.As(err, &usecase.ContentModelUpdateFailedError{}):
			c.JSON(http.StatusConflict, err.Error())
		default:
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		return nil
	}

	resFields := make([]Field, len(contentModel.Fields))
	for i, field := range contentModel.Fields {
		resFields[i] = Field{
			Name:     field.Name.String(),
			Type:     field.Type.String(),
			Required: bool(field.Required),
		}
	}

	c.JSON(http.StatusOK, ContentModelResponseBody{
		ID:        contentModel.ID.String(),
		Name:      contentModel.Name.String(),
		Fields:    resFields,
		CreatedAt: contentModel.CreatedAt.Time().Format(time.RFC3339),
		UpdatedAt: contentModel.UpdatedAt.Time().Format(time.RFC3339),
	})

	return nil
}
