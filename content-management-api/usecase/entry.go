package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/port"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/util/log"
)

type Entry struct {
	EntryPort            port.Entry
	EntryPublicationPort port.EntryPublication
	EntryArchivePort     port.EntryArchive
	ContentModelPort     port.ContentModel
}

func NewEntry(
	entryPort port.Entry,
	publicationPort port.EntryPublication,
	archivePort port.EntryArchive,
	contentModelPort port.ContentModel,
) *Entry {
	return &Entry{
		EntryPort:            entryPort,
		EntryPublicationPort: publicationPort,
		EntryArchivePort:     archivePort,
		ContentModelPort:     contentModelPort,
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

	if err := entry.CompareToModel(contentModel); err != nil {
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

func (e *Entry) Find(contentModelId domain.ContentModelID) (domain.Entries, error) {
	entries, err := e.EntryPort.Find(context.TODO())

	filter := entries.Filter(func(entry domain.Entry) bool {
		if contentModelId.IsEmpty() {
			return true
		}
		return entry.ContentModelID.String() == contentModelId.String()
	})

	if err != nil {
		switch err.(type) {
		default:
			return nil, fmt.Errorf("reason: %w", err)
		}
	}

	return filter, nil
}

func (e *Entry) FindByID(id domain.EntryId) (domain.Entry, error) {

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

func (e *Entry) DeleteByID(id domain.EntryId) error {

	err := e.EntryPort.DeleteById(context.TODO(), id)

	if err != nil {
		switch err.(type) {
		case domain.EntryNotFound:
			return NewEntryNotFoundError(err.Error())
		default:
			return err
		}
	}

	return nil
}

func (e *Entry) Published(id domain.EntryId) error {
	ctx := context.TODO()

	entryPublication := domain.EntryPublication{
		EntryId:         id,
		PublishedStatus: true,
	}

	err := e.EntryPublicationPort.Store(ctx, entryPublication)

	if err != nil {
		switch err.(type) {
		case domain.EntryNotFound:
			return NewEntryNotFoundError(err.Error())
		default:
			return err
		}
	}

	return nil
}

func (e *Entry) UnPublished(id domain.EntryId) error {
	ctx := context.TODO()

	entryPublication := domain.EntryPublication{
		EntryId:         id,
		PublishedStatus: false,
	}

	err := e.EntryPublicationPort.Store(ctx, entryPublication)

	if err != nil {
		switch err.(type) {
		case domain.EntryNotFound:
			return NewEntryNotFoundError(err.Error())
		default:
			return err
		}
	}

	return nil
}

func (e *Entry) Archive(id domain.EntryId) error {
	ctx := context.TODO()

	entryArchive := domain.EntryArchive{
		EntryId:        id,
		ArchivedStatus: true,
	}

	e.EntryArchivePort.Store(ctx, entryArchive)

	return nil
}
