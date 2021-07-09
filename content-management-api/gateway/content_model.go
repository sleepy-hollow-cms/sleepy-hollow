package gateway

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/driver/mongo"
	"content-management-api/usecase/write"
	"context"
)

type ContentModel struct {
	Driver mongo.ContentModelDriverInterface
}

func NewContentModel(driver mongo.ContentModelDriverInterface) *ContentModel {
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

func (c *ContentModel) Save(ctx context.Context, contentModel domain.ContentModel) error {
	panic("implement me")
}

func (c *ContentModel) Create(ctx context.Context, contentModel write.ContentModel) (domain.ContentModel, error) {
	fields := make([]string, len(contentModel.Fields))

	for i, field := range contentModel.Fields {
		fields[i] = field.Type.String()
	}

	created, err := c.Driver.Create(fields)

	if err != nil {
		return domain.ContentModel{}, err
	}

	createdFields := make(field.Fields, len(created.Fields))
	for i, createdField := range created.Fields {
		createdFields[i] = field.Field{
			Type: field.Of(createdField.Type),
		}
	}

	return domain.ContentModel{
		ID:     domain.ContentModelID(created.ID.Hex()),
		Fields: createdFields,
	}, nil
}
