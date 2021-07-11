package port

import (
	"content-management-api/domain"
	"content-management-api/usecase/write"
	"context"
)

type ContentModel interface {
	FindByID(ctx context.Context, id domain.ContentModelID) (domain.ContentModel, error)
	FindBySpaceID(ctx context.Context, id domain.SpaceID) (domain.ContentModels, error)
	Create(ctx context.Context, contentModel write.ContentModel) (domain.ContentModel, error)
}