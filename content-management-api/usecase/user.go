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

func (u *User) DeleteById(id domain.UserId) error {

	err := u.UserPort.DeleteById(context.TODO(), id)

	if err != nil {
		switch err.(type) {
		case domain.UserNotFound:
			return NewUserNotFoundError(err.Error())
		default:
			return err
		}
	}

	return nil
}
