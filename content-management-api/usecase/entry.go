package usecase

import (
	"content-management-api/port"
	"content-management-api/usecase/read"
	"content-management-api/usecase/write"
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

func (e *Entry) Register(entry write.Entry, entryItems []write.EntryItem) (read.Entry, error) {
	todoCtx := context.TODO()

	if _, err := e.ContentModelPort.FindByID(todoCtx, entry.ContentModelID); err != nil {
		switch {
		case errors.As(err, &ContentModelNotFoundError{}):
			return read.Entry{}, err
		default:
			return read.Entry{}, err
		}
	}

	createdEntry, err := e.EntryPort.Create(todoCtx, entry)

	if err != nil {
		return read.Entry{}, err
	}

	createdEntryItems, err := e.EntryPort.CreateItems(todoCtx, createdEntry.ID, entryItems)

	if err != nil {
		log.Logger.Warn(err)
		return read.Entry{}, err
	}

	return read.Entry{
		ID:             createdEntry.ID,
		ContentModelID: createdEntry.ContentModelID,
		EntryItems:     createdEntryItems,
	}, nil
}
