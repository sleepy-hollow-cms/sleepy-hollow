package gateway

import (
	"context"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver"
)

type EntryArchive struct {
	Driver driver.ContentDriver
}

func NewEntryArchive(driver driver.ContentDriver) *EntryArchive {
	return &EntryArchive{
		Driver: driver,
	}
}

func (e *EntryArchive) Store(ctx context.Context, entryPublication domain.EntryArchive) error {
	panic("not impl")
	//entry := model.Entry{
	//	ID: entryPublication.EntryId.String(),
	//	Publication: model.Publication{
	//		Status: entryPublication.PublishedStatus,
	//	},
	//}
	//
	//_, err := e.Driver.UpdateEntry(entry)
	//
	//return err
}

func (e *EntryArchive) Delete(ctx context.Context, entryPublication domain.EntryArchive) error {
	panic("not impl")
}
