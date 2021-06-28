package gateway

import (
	"content-management-api/domain"
	"context"
)

type Space struct{}

func NewSpace() *Space {
	return &Space{}
}

func (c *Space) FindByID(ctx context.Context, id domain.SpaceID) (domain.Space, error) {
	panic("implement me")
}
