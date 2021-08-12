package gateway

import (
	"content-management-api/domain"
	"content-management-api/driver"
	"content-management-api/driver/model"
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

func (e *Entry) Create(ctx context.Context, entry domain.Entry) (domain.Entry, error) {

	create, err := e.Driver.CreateEntry(model.Entry{ModelID: entry.ContentModelID.String()})

	if err != nil {
		return domain.Entry{}, err
	}

	return domain.Entry{
		ID:             domain.EntryId(create.ID),
		ContentModelID: domain.ContentModelID(create.ModelID),
	}, err
}

func (e *Entry) CreateItems(ctx context.Context, id domain.EntryId, items []domain.EntryItem) ([]domain.EntryItem, error) {
	entryItems := make([]model.EntryItem, len(items))
	for i, item := range items {
		entryItems[i] = model.EntryItem{
			Type:  item.Type.String(),
			Name:  item.FieldName.String(),
			Value: item.Value.FieldValue(),
		}
	}

	_, err := e.Driver.CreateEntryItems(model.EntryID(id.String()), entryItems)

	if err != nil {
		return make([]domain.EntryItem, 0), err
	}

	return items, nil
}
