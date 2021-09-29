package port

import (
	"context"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
)

type EntryPublication interface {
	Store(ctx context.Context, entryPublication domain.EntryPublication) error
	Delete(ctx context.Context, entryPublication domain.EntryPublication) error
}
