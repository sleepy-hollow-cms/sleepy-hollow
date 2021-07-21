package driver

import "content-management-api/driver/model"

// ContentModel is data-store driver
type ContentModel interface {
	Create(string, []model.Field) (*model.ContentModel, error)
}

type Entry interface {
	Create(model.Entry) (*model.Entry, error)
}
