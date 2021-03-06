package driver

import (
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
)

type UserDriver interface {
	Register(user model.User) (model.User, error)
	DeleteById(id string) (int64, error)
	FindById(id string) (*model.User, error)
	Update(user model.User) (*model.User, error)
}
