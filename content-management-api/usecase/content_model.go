package usecase

import (
	"content-management-api/domain"
	"content-management-api/port"
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

func (c *ContentModel) CreateContentModel(contentModel domain.ContentModel) (domain.ContentModel, error) {
	err := c.ContentModelPort.Save(context.TODO(), contentModel)

	if err != nil {
		return domain.ContentModel{}, NewContentModelSaveFailError(fmt.Sprintf("content model: %+v", contentModel))
	}

	return contentModel, nil
}

func (c *ContentModel) FindContentModelByID(id domain.ContentModelID) (domain.ContentModel, error) {
	result, err := c.ContentModelPort.FindByID(context.TODO(), id)

	if err != nil {
		return domain.ContentModel{}, NewContentModelNotFoundError(fmt.Sprintf("content model ID: %+v", id))
	}

	return result, nil
}

func (c *ContentModel) FindContentModelBySpaceID(spaceID domain.SpaceID) (domain.ContentModels, error) {
	return c.ContentModelPort.FindBySpaceID(context.TODO(), spaceID)
}
