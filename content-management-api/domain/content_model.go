package domain

type ContentModelID string

func (c ContentModelID) String() string {
	return string(c)
}

type ContentModel struct {
	ID ContentModelID
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
