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
	modelSpaces, err := s.Driver.FindSpace()
	if err != nil {
		return nil, err
	}
	spaces := make(domain.Spaces, len(modelSpaces))
	for i, modelSpace := range modelSpaces {
		spaces[i] = domain.Space{
			ID:   domain.SpaceID(modelSpace.ID),
			Name: domain.Name(modelSpace.Name),
		}
	}
	return spaces, nil
}

func (s *Space) FindByID(ctx context.Context, id domain.SpaceID) (domain.Space, error) {
	found, err := s.Driver.FindSpaceByID(id.String())

	if err != nil {
		return domain.Space{}, err
	}

	return domain.Space{
		ID:   domain.SpaceID(found.ID),
		Name: domain.Name(found.Name),
	}, nil
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
