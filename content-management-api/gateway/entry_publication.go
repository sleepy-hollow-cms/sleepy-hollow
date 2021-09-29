package gateway

import (
	"context"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver"
)

type EntryPublication struct {
	Driver driver.ContentDriver
}

func NewEntryPublication(driver driver.ContentDriver) *Entry {
	return &Entry{
		Driver: driver,
	}
}

func (e *Entry) Store(ctx context.Context, entryPublication domain.EntryPublication) error {
	panic("not impl")
}

func (e *Entry) Delete(ctx context.Context, entryPublication domain.EntryPublication) error {
	panic("not impl")
}
