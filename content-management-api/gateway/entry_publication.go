package gateway

import (
	"context"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver"
)

type EntryPublication struct {
	Driver driver.ContentDriver
}

func NewEntryPublication(driver driver.ContentDriver) *EntryPublication {
	return &EntryPublication{
		Driver: driver,
	}
}

func (e *EntryPublication) Store(ctx context.Context, entryPublication domain.EntryPublication) error {

	entry := model.Entry{
		ID: entryPublication.EntryId.String(),
		Publication: model.Publication{
			Status: entryPublication.PublishedStatus,
		},
	}

	_, err := e.Driver.UpdateEntry(entry)

	return err
}

func (e *EntryPublication) Delete(ctx context.Context, entryPublication domain.EntryPublication) error {
	panic("not impl")
}
