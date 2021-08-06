package usecase

import (
	"content-management-api/domain"
	"content-management-api/port"
	"content-management-api/usecase/write"
	"context"
	"errors"
	"fmt"
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

func (e *Entry) Register(entry write.Entry, entryItems []write.EntryItem) (domain.Entry, error) {
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
	createdEntryItems, err := e.EntryPort.CreateItems(todoCtx, entryItems)

	fmt.Println(createdEntryItems)

	if err != nil {
		return domain.Entry{}, err
	}

	return createdEntry, nil
}
