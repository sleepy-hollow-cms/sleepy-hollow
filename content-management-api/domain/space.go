package domain

type SpaceID string

func (i SpaceID) String() string {
	return string(i)
}

type Space struct {
	ID SpaceID
}
