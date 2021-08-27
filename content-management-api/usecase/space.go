package usecase

import (
	"content-management-api/domain"
	"content-management-api/port"
	"context"
)

type Space struct {
	SpacePort port.Space
}

func NewSpace(spacePort port.Space) *Space {
	return &Space{
		SpacePort: spacePort,
	}
}

func (s *Space) Find() (domain.Spaces, error) {
	spaces, err := s.SpacePort.Find(context.TODO())

	if err != nil {
		return nil, err
	}

	return spaces, nil
}

func (s *Space) FindByID(id domain.SpaceID) (domain.Space, error) {

	space, err := s.SpacePort.FindByID(context.TODO(), id)

	if err != nil {
		return domain.Space{}, NewSpaceNotFoundError("")
	}

	return space, nil
}

func (s *Space) Register(space domain.Space) (domain.Space, error) {
	space, err := s.SpacePort.Register(context.TODO(), space)

	if err != nil {
		return domain.Space{}, NewSpaceCreateFailedError(err.Error())
	}

	return space, nil
}

func (s *Space) Update(space domain.Space) (domain.Space, error) {

	found, err := s.SpacePort.FindByID(context.TODO(), space.ID)

	if err != nil {
		return domain.Space{}, NewSpaceNotFoundError(err.Error())
	}

	updatedSpace, err := s.SpacePort.Update(context.TODO(), domain.Space{
		ID:        space.ID,
		Name:      space.Name,
		CreatedAt: found.CreatedAt,
		UpdatedAt: space.UpdatedAt,
	})

	if err != nil {
		return domain.Space{}, NewSpaceUpdateFailedError(err.Error())
	}

	return updatedSpace, nil
}
