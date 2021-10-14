package port

import (
"context"
"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
)

type User interface {
	Register(ctx context.Context, user domain.User) (domain.User, error)
}
