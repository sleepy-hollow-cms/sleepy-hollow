package port

import (
	"content-management-api/domain"
	"context"
)

type ContentModel interface {
	FindByID(ctx context.Context, id domain.ContentModelID) (domain.ContentModel, error)
	FindBySpaceID(ctx context.Context, id domain.SpaceID) (domain.ContentModels, error)
	Save(ctx context.Context, contentModel domain.ContentModel) error
}
