package driver

import "content-management-api/driver/model"

// ContentModel is data-store driver
type ContentModel interface {
	Create(string, []string) (*model.ContentModel, error)
}
