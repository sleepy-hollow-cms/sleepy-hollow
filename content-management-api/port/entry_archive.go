package port

import (
	"context"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
)

type EntryArchive interface {
	Store(ctx context.Context, entryArchive domain.EntryArchive) error
	Delete(ctx context.Context, entryArchive domain.EntryArchive) error
}
