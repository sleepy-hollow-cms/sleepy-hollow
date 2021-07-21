package port

import (
	"content-management-api/domain"
	"context"
)

type Entry interface {
	Create(ctx context.Context) (domain.Entry, error)
}
