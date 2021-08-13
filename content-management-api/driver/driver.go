package driver

import (
	"content-management-api/driver/model"
	"time"
)

// ContentDriver is data-store driver
type ContentDriver interface {
	// CreateModel ContentModel driver interfaces
	CreateModel(string, time.Time, []model.Field) (*model.ContentModel, error)
	UpdateModel(model.ContentModel) (*model.ContentModel, error)
	FindContentModelByID(string) (*model.ContentModel, error)
	FindContentModelBySpaceID(string) ([]model.ContentModel, error)
	DeleteContentModelByID(string) error

	CreateEntry(model.Entry) (*model.Entry, error)
	CreateEntryItems(model.EntryID, []model.EntryItem) ([]model.EntryItem, error)
	FindEntryByID(string) (*model.Entry, error)
}
