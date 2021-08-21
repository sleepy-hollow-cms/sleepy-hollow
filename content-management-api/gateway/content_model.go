package gateway

import (
	"content-management-api/domain"
	"content-management-api/driver"
	"content-management-api/driver/model"
	"content-management-api/usecase"
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
		UpdatedAt: domain.UpdatedAt(contentModels.UpdatedAt),
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
			UpdatedAt: domain.UpdatedAt(foundModel.UpdatedAt),
			Fields:    newFields(foundModel.Fields),
		}
	}

	return domain.ContentModels{
		SpaceID: id,
		Models:  foundModels,
	}, nil
}

func (c *ContentModel) Create(ctx context.Context, contentModel domain.ContentModel) (domain.ContentModel, error) {
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
		UpdatedAt: domain.UpdatedAt(created.UpdatedAt),
	}, nil
}

func (c *ContentModel) Update(ctx context.Context, foundContentModel domain.ContentModel, updatedContentModel domain.ContentModel) (domain.ContentModel, error) {
	if foundContentModel.UpdatedAt != updatedContentModel.UpdatedAt {
		return domain.ContentModel{}, usecase.NewContentModelUpdateFailedError("Content Model Update conflicted")
	}

	fields := make([]model.Field, len(foundContentModel.Fields))
	for i, field := range foundContentModel.Fields {
		fields[i] = model.Field{
			Name:     field.Name.String(),
			Type:     field.Type.String(),
			Required: bool(field.Required),
		}
	}

	updated, err := c.Driver.UpdateModel(
		model.ContentModel{
			ID:        foundContentModel.ID.String(),
			Name:      foundContentModel.Name.String(),
			CreatedAt: foundContentModel.CreatedAt.Time(),
			UpdatedAt: foundContentModel.UpdatedAt.Time(),
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
		UpdatedAt: domain.UpdatedAt(updated.UpdatedAt),
	}, nil
}

func newFields(modelFields []model.Field) []domain.Field {
	fields := make(domain.Fields, len(modelFields))
	for i, getField := range modelFields {
		fields[i] = domain.Field{
			Name:     domain.Name(getField.Name),
			Type:     domain.Of(getField.Type),
			Required: domain.Required(getField.Required),
		}
	}
	return fields
}
