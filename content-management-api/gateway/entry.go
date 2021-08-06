package gateway

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/driver"
	"content-management-api/driver/model"
	"content-management-api/usecase/write"
	"content-management-api/util/log"
	"context"
)

type Entry struct {
	Driver driver.ContentDriver
}

func NewEntry(driver driver.ContentDriver) *Entry {
	return &Entry{
		Driver: driver,
	}
}

func (e *Entry) Create(ctx context.Context, entry write.Entry) (domain.Entry, error) {

	create, err := e.Driver.CreateEntry(model.Entry{ModelID: entry.ContentModelID.String()})

	if err != nil {
		return domain.Entry{}, err
	}

	return domain.Entry{
		ID:             domain.EntryId(create.ID),
		ContentModelID: domain.ContentModelID(create.ModelID),
	}, err
}

func (e *Entry) CreateItems(ctx context.Context, items []write.EntryItem) ([]field.HasValue, error) {
	log.Logger.Warn("not implemented")
	return nil, nil
}
