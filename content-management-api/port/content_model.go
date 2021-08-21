package port

import (
	"content-management-api/domain"
	"context"
)

type ContentModel interface {
	FindByID(ctx context.Context, id domain.ContentModelID) (domain.ContentModel, error)
	DeleteByID(ctx context.Context, id domain.ContentModelID) error
	FindBySpaceID(ctx context.Context, id domain.SpaceID) (domain.ContentModels, error)
	Create(ctx context.Context, contentModel domain.ContentModel) (domain.ContentModel, error)
	Update(ctx context.Context, foundContentModel domain.ContentModel, updatedContentModel domain.ContentModel) (domain.ContentModel, error)
}
