package usecase

import (
	"content-management-api/domain"
	"content-management-api/port"
	"content-management-api/util/log"
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

func (e *Entry) Register(entry domain.Entry) (domain.Entry, error) {
	todoCtx := context.TODO()

	contentModel, err := e.ContentModelPort.FindByID(todoCtx, entry.ContentModelID)
	if err != nil {
		switch {
		case errors.As(err, &ContentModelNotFoundError{}):
			return domain.Entry{}, err
		default:
			return domain.Entry{}, err
		}
	}

	if err := entry.Validate(contentModel); err != nil {
		return domain.Entry{}, err
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

func (e *Entry) Find(id domain.EntryId) (domain.Entry, error) {

	result, err := e.EntryPort.FindById(context.TODO(), id)

	if err != nil {
		switch err.(type) {
		case domain.EntryNotFound:
			return domain.Entry{}, NewEntryNotFoundError(err.Error())
		default:
			return domain.Entry{}, fmt.Errorf("reason: %w", err)
		}
	}

	return result, nil
}
