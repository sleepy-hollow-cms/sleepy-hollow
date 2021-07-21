package gateway

import (
	"content-management-api/domain"
	"content-management-api/driver"
	"content-management-api/driver/model"
	"content-management-api/usecase/write"
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

func (e *Entry) Create(ctx context.Context, entry write.Entry) (domain.Entry, error) {

	create, err := e.Driver.Create(model.Entry{ModelID: entry.ContentModelID.String()})

	if err != nil {
		return domain.Entry{}, err
	}

	return domain.Entry{
		ID:             domain.EntryId(create.ID),
		ContentModelID: domain.ContentModelID(create.ModelID),
	}, err
}
