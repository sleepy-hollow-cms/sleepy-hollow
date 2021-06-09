package domain

type ContentModelID string

type ContentModel struct {
	ID ContentModelID
}

type ContentModels struct {
	list []ContentModel
}
