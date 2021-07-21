package usecase

import (
	"content-management-api/domain"
	"content-management-api/port"
	"content-management-api/usecase/write"
	"context"
)

type Entry struct {
	EntryPort        port.Entry
	ContentModelPort port.ContentModel
}

func NewEntry(
	entryPort port.Entry,
) *Entry {
	return &Entry{
		EntryPort: entryPort,
	}
}

func (e *Entry) Create(entry write.Entry) (domain.Entry, error) {
	return e.EntryPort.Create(context.TODO(), entry)
}
