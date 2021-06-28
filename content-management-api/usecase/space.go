package usecase

import (
	"content-management-api/domain"
	"content-management-api/port"
)

type Space struct {
	SpacePort port.Space
}

func NewSpace(spacePort port.Space) *Space {
	return &Space{
		SpacePort: spacePort,
	}
}

func (s *Space) FindByID(id domain.SpaceID) (domain.Space, error) {

	space, err := s.SpacePort.FindByID(id)

	if err != nil {
		return domain.Space{}, NewSpaceNotFoundError("")
	}

	return space, nil
}
