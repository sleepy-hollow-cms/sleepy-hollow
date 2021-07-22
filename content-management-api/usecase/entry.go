package usecase

import (
	"content-management-api/domain"
	"content-management-api/port"
	"content-management-api/usecase/write"
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

func (e *Entry) Create(entry write.Entry) (domain.Entry, error) {
	if _, err := e.ContentModelPort.FindByID(context.TODO(), entry.ContentModelID); err != nil {
		switch {
		case errors.As(err, &ContentModelNotFoundError{}):
			return domain.Entry{}, err
		default:
			return domain.Entry{}, err
		}
	}
	create, err := e.EntryPort.Create(context.TODO(), entry)
	if err != nil {
		return domain.Entry{}, err
	}
	return create, nil
}
