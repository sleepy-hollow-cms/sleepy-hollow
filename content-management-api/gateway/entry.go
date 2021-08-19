package gateway

import (
	"content-management-api/domain"
	"content-management-api/driver"
	"content-management-api/driver/model"
	"content-management-api/util"
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

func (e *Entry) FindById(ctx context.Context, entryId domain.EntryId) (domain.Entry, error) {

	found, err := e.Driver.FindEntryByID(entryId.String())

	if err != nil {
		return domain.Entry{}, err
	}

	items := make([]domain.EntryItem, len(found.Items))
	errs := new(util.ErrorCollector)
	for i, item := range found.Items {
		v, err := domain.FactoryValue(item.Value)

		if err != nil {
			errs.Collect(err)
		}

		items[i] = domain.EntryItem{
			Value: v,
		}
	}

	return domain.Entry{
		ID:             domain.EntryId(found.ID),
		ContentModelID: domain.ContentModelID(found.ModelID),
		Items:          items,
	}, nil
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
			Value: item.Value,
		}
	}

	_, err := e.Driver.CreateEntryItems(model.EntryID(id.String()), entryItems)

	if err != nil {
		return make([]domain.EntryItem, 0), err
	}

	return items, nil
}
