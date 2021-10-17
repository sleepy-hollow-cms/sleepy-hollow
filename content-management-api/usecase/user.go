package usecase

import (
	"context"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/port"
)

type User struct {
	UserPort port.User
}

func NewUser(userPort port.User) *User {
	return &User{
		UserPort: userPort,
	}
}

func (u *User) Register(user domain.User) (domain.User, error) {
	user, err := u.UserPort.Register(context.TODO(), user)
	if err != nil {
		return domain.User{}, NewUserCreateFailedError(err.Error())
	}
	return user, nil
}
