package driver

import (
	"time"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
)

// ContentDriver is data-store driver
type ContentDriver interface {
	FindSpace() ([]model.Space, error)
	FindSpaceByID(id string) (*model.Space, error)
	CreateSpace(space model.Space) (*model.Space, error)
	UpdateSpace(space model.Space) (*model.Space, error)
	// CreateModel ContentModel driver interfaces
	CreateModel(string, time.Time, []model.Field) (*model.ContentModel, error)
	UpdateModel(model.ContentModel) (*model.ContentModel, error)
	FindContentModelByID(string) (*model.ContentModel, error)
	FindContentModelBySpaceID(string) ([]model.ContentModel, error)
	DeleteContentModelByID(string) error
	CreateEntry(model.Entry) (*model.Entry, error)
	CreateEntryItems(model.EntryID, []model.EntryItem) ([]model.EntryItem, error)
	FindEntry() ([]model.Entry, error)
	FindEntryByID(string) (*model.Entry, error)
	DeleteEntryByID(string) (int64, error)
}
