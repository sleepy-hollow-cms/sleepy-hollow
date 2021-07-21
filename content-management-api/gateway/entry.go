package gateway

import (
	"content-management-api/domain"
	"content-management-api/driver"
	"context"
)

type Entry struct {
	Driver driver.Entry
}

func NewEntry(driver driver.Entry) *Entry {
	return &Entry{
		Driver: driver,
	}
}

func (e *Entry) Create(ctx context.Context) (domain.Entry, error) {
	create, err := e.Driver.Create()

	if err != nil {
		return domain.Entry{}, err
	}

	return domain.Entry{ID: domain.EntryId(create.ID)}, err
}
