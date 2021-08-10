package gateway

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/driver"
	"content-management-api/driver/model"
	"content-management-api/usecase/read"
	"content-management-api/usecase/write"
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

func (e *Entry) CreateItems(ctx context.Context, id domain.EntryId, items []write.EntryItem) (read.EntryItem, error) {
	entryItems := make([]model.EntryItem, len(items))
	for i, item := range items {
		entryItems[i] = model.EntryItem{
			Type:  item.Type.String(),
			Name:  item.FieldName.String(),
			Value: item.Value.FieldValue(),
		}
	}

	createEntryItems, err := e.Driver.CreateEntryItems(model.EntryID(id.String()), entryItems)

	if err != nil {
		return read.EntryItem{}, err
	}

	readItems := make([]read.Item, len(createEntryItems))
	for i, createEntryItem := range createEntryItems {
		itemType := field.Of(createEntryItem.Type)
		readItems[i] = read.Item{
			FieldName: field.Name(createEntryItem.Name),
			Type:      itemType,
			Value:     field.FactoryValue(itemType, createEntryItem.Value),
		}
	}

	return read.EntryItem{
		ID:    field.ID(id),
		Items: readItems,
	}, nil
}
