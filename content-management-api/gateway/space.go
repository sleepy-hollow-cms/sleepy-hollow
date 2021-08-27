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

func (s *Space) Find(ctx context.Context) (domain.Spaces, error) {
	driverResults, err := s.Driver.FindSpace()
	if err != nil {
		return nil, err
	}
	spaces := make(domain.Spaces, len(driverResults))
	for i, driverResult := range driverResults {
		spaces[i] = domain.NewSpace(
			domain.SpaceID(driverResult.ID),
			domain.Name(driverResult.Name),
			domain.CreatedAt(driverResult.CreatedAt),
			domain.UpdatedAt(driverResult.UpdatedAt),
		)
	}
	return spaces, nil
}

func (s *Space) FindByID(ctx context.Context, id domain.SpaceID) (domain.Space, error) {
	driverResult, err := s.Driver.FindSpaceByID(id.String())

	if err != nil {
		return domain.Space{}, err
	}

	return domain.NewSpace(
		domain.SpaceID(driverResult.ID),
		domain.Name(driverResult.Name),
		domain.CreatedAt(driverResult.CreatedAt),
		domain.UpdatedAt(driverResult.UpdatedAt),
	), nil
}

func (s *Space) Register(ctx context.Context, space domain.Space) (domain.Space, error) {
	driverResult, err := s.Driver.CreateSpace(model.Space{
		Name:      space.Name.String(),
		CreatedAt: space.CreatedAt.Time(),
		UpdatedAt: space.UpdatedAt.Time(),
	})

	if err != nil {
		return domain.Space{}, nil
	}

	return domain.NewSpace(
		domain.SpaceID(driverResult.ID),
		domain.Name(driverResult.Name),
		domain.CreatedAt(driverResult.CreatedAt),
		domain.UpdatedAt(driverResult.UpdatedAt),
	), nil
}

func (s *Space) Update(ctx context.Context, space domain.Space) (domain.Space, error) {
	driverResult, err := s.Driver.UpdateSpace(model.Space{
		Name:      space.Name.String(),
		CreatedAt: space.CreatedAt.Time(),
		UpdatedAt: space.UpdatedAt.Time(),
	})

	if err != nil {
		return domain.Space{}, nil
	}

	return domain.NewSpace(
		domain.SpaceID(driverResult.ID),
		domain.Name(driverResult.Name),
		domain.CreatedAt(driverResult.CreatedAt),
		domain.UpdatedAt(driverResult.UpdatedAt),
	), nil
}
