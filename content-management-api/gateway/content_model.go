package gateway

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/driver"
	"content-management-api/driver/model"
	"content-management-api/usecase/write"
	"context"
)

type ContentModel struct {
	Driver driver.ContentModel
}

func NewContentModel(driver driver.ContentModel) *ContentModel {
	return &ContentModel{
		Driver: driver,
	}
}

func (c *ContentModel) FindByID(ctx context.Context, id domain.ContentModelID) (domain.ContentModel, error) {
	panic("implement me")
}

func (c *ContentModel) FindBySpaceID(ctx context.Context, id domain.SpaceID) (domain.ContentModels, error) {
	panic("implement me")
}

func (c *ContentModel) Create(ctx context.Context, contentModel write.ContentModel) (domain.ContentModel, error) {
	fields := make([]model.Field, len(contentModel.Fields))

	for i, field := range contentModel.Fields {
		fields[i] = model.Field{
			Name:     field.Name.String(),
			Type:     field.Type.String(),
			Required: bool(field.Required),
		}
	}

	created, err := c.Driver.Create(contentModel.Name.String(), fields)

	if err != nil {
		return domain.ContentModel{}, err
	}

	createdFields := make(field.Fields, len(created.Fields))
	for i, createdField := range created.Fields {
		createdFields[i] = field.Field{
			Name:     field.Name(createdField.Name),
			Type:     field.Of(createdField.Type),
			Required: field.Required(createdField.Required),
		}
	}

	return domain.ContentModel{
		ID:     domain.ContentModelID(created.ID),
		Name:   domain.Name(created.Name),
		Fields: createdFields,
	}, nil
}
