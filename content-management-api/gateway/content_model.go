package gateway

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/driver"
	"content-management-api/driver/model"
	"content-management-api/usecase"
	"content-management-api/usecase/write"
	"content-management-api/util/log"
	"context"
)

type ContentModel struct {
	Driver driver.ContentDriver
}

func NewContentModel(driver driver.ContentDriver) *ContentModel {
	return &ContentModel{
		Driver: driver,
	}
}

func (c *ContentModel) FindByID(ctx context.Context, id domain.ContentModelID) (domain.ContentModel, error) {

	contentModels, err := c.Driver.FindContentModelByID(id.String())

	if err != nil {
		switch err := err.(type) {
		case driver.ContentModelCannotFindByIdError:
			log.Logger.Warn(err.Error())
			return domain.ContentModel{}, usecase.NewContentModelNotFoundError(err.Error())
		default:
			log.Logger.Warn(err.Error())
			return domain.ContentModel{}, err
		}
	}
	fields := make(field.Fields, len(contentModels.Fields))

	for i, getField := range contentModels.Fields {
		fields[i] = field.Field{
			Name:     field.Name(getField.Name),
			Type:     field.Of(getField.Type),
			Required: field.Required(getField.Required),
		}
	}

	return domain.ContentModel{
		ID:     domain.ContentModelID(contentModels.ID),
		Name:   domain.Name(contentModels.Name),
		Fields: fields,
	}, nil

}

func (c *ContentModel) DeleteByID(ctx context.Context, id domain.ContentModelID) error {

	err := c.Driver.DeleteContentModelByID(id.String())

	if err != nil {
		return err
	}

	return nil

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

	created, err := c.Driver.CreateModel(contentModel.Name.String(), fields)

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
