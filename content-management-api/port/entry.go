package port

import (
	"content-management-api/domain"
	"content-management-api/usecase/read"
	"content-management-api/usecase/write"
	"context"
)

type Entry interface {
	Create(ctx context.Context, entry write.Entry) (domain.Entry, error)
	CreateItems(ctx context.Context, entryId domain.EntryId, entry []write.EntryItem) (read.EntryItem, error)
}
