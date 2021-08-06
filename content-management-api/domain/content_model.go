package domain

import (
	"content-management-api/domain/field"
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

func (c ContentModelID) String() string {
	return string(c)
}

type ContentModel struct {
	ID        ContentModelID
	Name      Name
	Fields    field.FieldModels
	CreatedAt CreatedAt
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
