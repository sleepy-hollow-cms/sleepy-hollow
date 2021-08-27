package port

import (
	"context"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
)

type Space interface {
	Find(ctx context.Context) (domain.Spaces, error)
	FindByID(ctx context.Context, id domain.SpaceID) (domain.Space, error)
	Register(ctx context.Context, space domain.Space) (domain.Space, error)
	Update(ctx context.Context, space domain.Space) (domain.Space, error)
}
