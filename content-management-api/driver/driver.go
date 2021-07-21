package driver

import "content-management-api/driver/model"

// ContentDriver is data-store driver
type ContentDriver interface {
	CreateModel(string, []model.Field) (*model.ContentModel, error)
	CreateEntry(model.Entry) (*model.Entry, error)
	FindContentModelByID(string) (*model.ContentModel, error)
}
