package usecase

import (
	"content-management-api/domain"
	"content-management-api/port"
	"content-management-api/support"
	"context"
	"fmt"
)

type ContentModelUseCase struct {
	ContentModelPort port.ContentModel
	SpacePort        port.Space
}

func (c *ContentModelUseCase) CreateContentModel(contentModel domain.ContentModel) (domain.ContentModel, error) {

	err := c.ContentModelPort.Save(context.TODO(), contentModel)

	if err != nil {
		return domain.ContentModel{}, NewContentModelSaveFailError(fmt.Sprintf("content model: %+v", contentModel))
	}

	return contentModel, nil
}

func (c *ContentModelUseCase) FindContentModelByID(id domain.ContentModelID) (domain.ContentModel, error) {
	result, err := c.ContentModelPort.FindByID(context.TODO(), id)

	if err != nil {
		return domain.ContentModel{}, NewContentModelNotFoundError(fmt.Sprintf("content model ID: %+v", id))
	}

	return result, nil
}

func (c ContentModelUseCase) FindContentModelBySpaceID(id domain.SpaceID) (*domain.ContentModels, error) {
	return nil, support.TODO("")
}
