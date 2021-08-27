package domain

type SpaceID string

func (i SpaceID) String() string {
	return string(i)
}

type Space struct {
	ID        SpaceID
	Name      Name
	CreatedAt CreatedAt
	UpdatedAt UpdatedAt
}

func NewSpace(id SpaceID, name Name, createdAt CreatedAt, updatedAt UpdatedAt) Space {
	return Space{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type Spaces []Space
