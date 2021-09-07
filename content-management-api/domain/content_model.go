package domain

import (
	"time"
)

type ContentModelID string
type Name string

func (n Name) String() string {
	return string(n)
}

type CreatedAt time.Time

func (d CreatedAt) Time() time.Time {
	return time.Time(d)
}

type UpdatedAt time.Time

func (d UpdatedAt) Time() time.Time {
	return time.Time(d)
}

func (c ContentModelID) String() string {
	return string(c)
}

func (c ContentModelID) IsEmpty() bool {
	return IsEmpty(c.String())
}

func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return false
}

type ContentModel struct {
	ID        ContentModelID
	Name      Name
	Fields    Fields
	CreatedAt CreatedAt
	UpdatedAt UpdatedAt
}

type ContentModels struct {
	SpaceID SpaceID
	Models  []ContentModel
}

func NewContentModels(list []ContentModel) ContentModels {
	contentModels := make([]ContentModel, len(list))
	for i, m := range list {
		contentModels[i] = ContentModel{
			ID: m.ID,
		}
	}
	return ContentModels{
		Models: contentModels,
	}
}
