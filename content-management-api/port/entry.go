package port

import (
	"context"

	"github.com/sleepy-hollow-cms/content-management-api/domain"
)

type Entry interface {
	Create(ctx context.Context, entry domain.Entry) (domain.Entry, error)
	CreateItems(ctx context.Context, entryId domain.EntryId, entry []domain.EntryItem) ([]domain.EntryItem, error)
	FindById(ctx context.Context, entryId domain.EntryId) (domain.Entry, error)
}
