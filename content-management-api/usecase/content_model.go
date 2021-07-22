package usecase

import (
	"content-management-api/domain"
	"content-management-api/port"
	"content-management-api/usecase/write"
	"context"
	"fmt"
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

func (c *ContentModel) FindContentModelByID(id domain.ContentModelID) (domain.ContentModel, error) {
	result, err := c.ContentModelPort.FindByID(context.TODO(), id)

	if err != nil {
		return domain.ContentModel{}, NewContentModelNotFoundError(fmt.Sprintf("content model ID: %+v", id))
	}

	return result, nil
}

func (c *ContentModel) DeleteContentModelByID(id domain.ContentModelID) (error) {
	err := c.ContentModelPort.DeleteByID(context.TODO(), id)

	if err != nil {
		return NewContentModelNotFoundError(fmt.Sprintf("content model ID: %+v", id))
	}

	return nil
}

func (c *ContentModel) FindContentModelBySpaceID(spaceID domain.SpaceID) (domain.ContentModels, error) {
	return c.ContentModelPort.FindBySpaceID(context.TODO(), spaceID)
}

func (c *ContentModel) Create(contentModel write.ContentModel) (domain.ContentModel, error) {
	result, err := c.ContentModelPort.Create(context.TODO(), contentModel)

	if err != nil {
		return domain.ContentModel{}, NewContentModelCreateFailedError(err.Error())
	}

	return result, err
}
