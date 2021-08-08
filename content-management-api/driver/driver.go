package driver

import (
	"content-management-api/driver/model"
	"time"
)

// ContentDriver is data-store driver
type ContentDriver interface {
	CreateModel(string, time.Time, []model.Field) (*model.ContentModel, error)
	UpdateModel(model.ContentModel) (*model.ContentModel, error)
	CreateEntry(model.Entry) (*model.Entry, error)
	FindContentModelByID(string) (*model.ContentModel, error)
	FindContentModelBySpaceID(string) ([]model.ContentModel, error)
	DeleteContentModelByID(string) error
}
