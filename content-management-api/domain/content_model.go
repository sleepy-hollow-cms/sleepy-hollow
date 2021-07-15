package domain

import (
	"content-management-api/domain/field"
)

type ContentModelID string
type Name string

func (n Name) String() string {
	return string(n)
}

func (c ContentModelID) String() string {
	return string(c)
}

type ContentModel struct {
	ID     ContentModelID
	Name   Name
	Fields field.Fields
}

type ContentModels []ContentModel

func NewContentModels(list []ContentModel) ContentModels {
	contentModels := make(ContentModels, len(list))
	for i, m := range list {
		contentModels[i] = ContentModel{
			ID: m.ID,
		}
	}
	return contentModels
}
