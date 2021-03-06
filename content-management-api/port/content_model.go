package port

import (
	"context"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
)

type ContentModel interface {
	FindByID(ctx context.Context, id domain.ContentModelID) (domain.ContentModel, error)
	DeleteByID(ctx context.Context, id domain.ContentModelID) error
	FindBySpaceID(ctx context.Context, id domain.SpaceID) (domain.ContentModels, error)
	Create(ctx context.Context, contentModel domain.ContentModel) (domain.ContentModel, error)
	Update(ctx context.Context, updatedContentModel domain.ContentModel) (domain.ContentModel, error)
}
