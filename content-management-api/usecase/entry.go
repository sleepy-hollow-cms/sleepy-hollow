package usecase

import (
	"content-management-api/domain"
	"content-management-api/port"
	"content-management-api/util/log"
	"context"
	"errors"
)

type Entry struct {
	EntryPort        port.Entry
	ContentModelPort port.ContentModel
}

func NewEntry(
	entryPort port.Entry,
	contentModelPort port.ContentModel,
) *Entry {
	return &Entry{
		EntryPort:        entryPort,
		ContentModelPort: contentModelPort,
	}
}

func (e *Entry) Register(entry domain.Entry) (domain.Entry, error) {
	todoCtx := context.TODO()

	if _, err := e.ContentModelPort.FindByID(todoCtx, entry.ContentModelID); err != nil {
		switch {
		case errors.As(err, &ContentModelNotFoundError{}):
			return domain.Entry{}, err
		default:
			return domain.Entry{}, err
		}
	}

	createdEntry, err := e.EntryPort.Create(todoCtx, entry)

	if err != nil {
		return domain.Entry{}, err
	}

	createdEntryItems, err := e.EntryPort.CreateItems(todoCtx, createdEntry.ID, entry.Items)

	if err != nil {
		log.Logger.Warn(err)
		return domain.Entry{}, err
	}

	return domain.Entry{
		ID:             createdEntry.ID,
		ContentModelID: createdEntry.ContentModelID,
		Items:          createdEntryItems,
	}, nil
}
