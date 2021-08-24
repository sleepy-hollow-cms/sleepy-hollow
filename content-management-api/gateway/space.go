package gateway

import (
	"content-management-api/domain"
	"content-management-api/driver"
	"content-management-api/driver/model"
	"context"
)

type Space struct {
	Driver driver.ContentDriver
}

func NewSpace(driver driver.ContentDriver) *Space {
	return &Space{
		Driver: driver,
	}
}

func (s *Space) FindByID(ctx context.Context, id domain.SpaceID) (domain.Space, error) {
	panic("implement me")
}

func (s *Space) Register(ctx context.Context, space domain.Space) (domain.Space, error) {
	createSpace, err := s.Driver.CreateSpace(model.Space{
		Name: space.Name.String(),
	})

	if err != nil {
		return domain.Space{}, nil
	}

	return domain.Space{
		ID:   domain.SpaceID(createSpace.ID),
		Name: domain.Name(createSpace.Name),
	}, nil
}
