package usecase

import (
	"content-management-api/domain"
	"content-management-api/support"
)

type ContentModelUseCase struct {
}

func (c *ContentModelUseCase) CreateContentModel(contentModel domain.ContentModel) (*domain.ContentModel, error) {
	return nil, support.TODO("")
}

func (c *ContentModelUseCase) FindContentModelByID(id domain.ContentModelID) (*domain.ContentModel, error) {
	return nil, support.TODO("")
}

func (c ContentModelUseCase) FindContentModelBySpaceID(id domain.SpaceID) (*domain.ContentModels, error) {
	return nil, support.TODO("")
}
