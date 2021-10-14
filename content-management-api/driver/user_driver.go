package driver

import "github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"

type UserDriver interface {
	Register(user model.User) (*model.User, error)
}
