package usecase

import (
	"content-management-api/domain"
	"content-management-api/port"
	"content-management-api/usecase/write"
	"content-management-api/util/log"
	"context"
)

type ContentModel struct {
	ContentModelPort port.ContentModel
}

func NewContentModel(
	contentModelPort port.ContentModel,
) *ContentModel {
	return &ContentModel{
		ContentModelPort: contentModelPort,
	}
}

func (c *ContentModel) FindByID(id domain.ContentModelID) (domain.ContentModel, error) {
	result, err := c.ContentModelPort.FindByID(context.TODO(), id)

	if err != nil {
		log.Logger.Warn(err.Error())
		return domain.ContentModel{}, err
	}

	return result, nil
}

func (c *ContentModel) DeleteContentModelByID(id domain.ContentModelID) error {
	err := c.ContentModelPort.DeleteByID(context.TODO(), id)

	if err != nil {
		log.Logger.Warn(err.Error())
		return err
	}

	return nil
}

func (c *ContentModel) FindContentModelBySpaceID(spaceID domain.SpaceID) (domain.ContentModels, error) {
	models, err := c.ContentModelPort.FindBySpaceID(context.TODO(), spaceID)
	if err != nil {
		log.Logger.Warn(err.Error())
		return domain.ContentModels{}, err
	}
	return models, nil
}

func (c *ContentModel) Create(contentModel write.ContentModel) (domain.ContentModel, error) {
	result, err := c.ContentModelPort.Create(context.TODO(), contentModel)

	if err != nil {
		return domain.ContentModel{}, NewContentModelCreateFailedError(err.Error())
	}

	return result, err
}
