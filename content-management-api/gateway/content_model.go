package gateway

import (
	"content-management-api/domain"
	"context"
)

type ContentModel struct{}

func NewContentModel() *ContentModel {
	return &ContentModel{}
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
