package port

import (
	"content-management-api/domain"
	"context"
)

type Entry interface {
	Create(ctx context.Context, entry domain.Entry) (domain.Entry, error)
	CreateItems(ctx context.Context, entryId domain.EntryId, entry []domain.EntryItem) ([]domain.EntryItem, error)
	FindById(ctx context.Context, entryId domain.EntryId) (domain.Entry, error)
}
