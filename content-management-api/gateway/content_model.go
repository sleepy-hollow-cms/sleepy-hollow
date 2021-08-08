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
	"errors"
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
		switch {
		case errors.As(err, &driver.ContentModelCannotFindByIdError{}):
			log.Logger.Warn(err.Error())
			return domain.ContentModel{}, usecase.NewContentModelNotFoundError(err.Error())
		default:
			log.Logger.Warn(err.Error())
			return domain.ContentModel{}, err
		}
	}

	return domain.ContentModel{
		ID:        domain.ContentModelID(contentModels.ID),
		Name:      domain.Name(contentModels.Name),
		CreatedAt: domain.CreatedAt(contentModels.CreatedAt),
		Fields:    newFields(contentModels.Fields),
	}, nil
}

func (c *ContentModel) DeleteByID(ctx context.Context, id domain.ContentModelID) error {

	err := c.Driver.DeleteContentModelByID(id.String())

	if err != nil {
		switch {
		case errors.As(err, &driver.ContentModelCannotFindByIdError{}):
			log.Logger.Warn(err.Error())
			return usecase.NewContentModelNotFoundError(err.Error())
		default:
			log.Logger.Warn(err.Error())
			return err
		}
	}

	return nil

}

func (c *ContentModel) FindBySpaceID(ctx context.Context, id domain.SpaceID) (domain.ContentModels, error) {
	contentModels, err := c.Driver.FindContentModelBySpaceID(id.String())

	if err != nil {
		log.Logger.Warn(err.Error())
		return domain.ContentModels{}, err
	}

	foundModels := make([]domain.ContentModel, len(contentModels))

	for i, foundModel := range contentModels {
		foundModels[i] = domain.ContentModel{
			ID:        domain.ContentModelID(foundModel.ID),
			Name:      domain.Name(foundModel.Name),
			CreatedAt: domain.CreatedAt(foundModel.CreatedAt),
			Fields:    newFields(foundModel.Fields),
		}
	}

	return domain.ContentModels{
		SpaceID: id,
		Models:  foundModels,
	}, nil
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

	created, err := c.Driver.CreateModel(contentModel.Name.String(), contentModel.CreatedAt.Time(), fields)

	if err != nil {
		log.Logger.Warn(err.Error())
		return domain.ContentModel{}, err
	}

	return domain.ContentModel{
		ID:        domain.ContentModelID(created.ID),
		Name:      domain.Name(created.Name),
		Fields:    newFields(created.Fields),
		CreatedAt: domain.CreatedAt(created.CreatedAt),
	}, nil
}

func (c *ContentModel) Update(ctx context.Context, id domain.ContentModelID, contentModel write.ContentModel) (domain.ContentModel, error) {
	fields := make([]model.Field, len(contentModel.Fields))
	for i, field := range contentModel.Fields {
		fields[i] = model.Field{
			Name:     field.Name.String(),
			Type:     field.Type.String(),
			Required: bool(field.Required),
		}
	}

	updated, err := c.Driver.UpdateModel(
		model.ContentModel{
			ID:        id.String(),
			Name:      contentModel.Name.String(),
			CreatedAt: contentModel.CreatedAt.Time(),
			Fields:    fields,
		})

	if err != nil {
		log.Logger.Warn(err.Error())
		return domain.ContentModel{}, err
	}

	return domain.ContentModel{
		ID:        domain.ContentModelID(updated.ID),
		Name:      domain.Name(updated.Name),
		Fields:    newFields(updated.Fields),
		CreatedAt: domain.CreatedAt(updated.CreatedAt),
	}, nil
}

func newFields(modelFields []model.Field) []field.Field {
	fields := make(field.Fields, len(modelFields))
	for i, getField := range modelFields {
		fields[i] = field.Field{
			Name:     field.Name(getField.Name),
			Type:     field.Of(getField.Type),
			Required: field.Required(getField.Required),
		}
	}
	return fields
}
