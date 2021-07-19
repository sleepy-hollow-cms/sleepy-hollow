package usecase

import (
	"content-management-api/domain"
	"content-management-api/port"
	"context"
)

type Entry struct {
	EntryPort port.Entry
}

func NewEntry(
	entryPort port.Entry,
) *Entry {
	return &Entry{
		EntryPort: entryPort,
	}
}

func (e *Entry) Create() (domain.Entry, error) {
	return e.EntryPort.Create(context.TODO())
}
