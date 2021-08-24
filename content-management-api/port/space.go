package port

import (
	"content-management-api/domain"
	"context"
)

type Space interface {
	FindByID(ctx context.Context, id domain.SpaceID) (domain.Space, error)
	Register(ctx context.Context, space domain.Space) (domain.Space, error)
}
