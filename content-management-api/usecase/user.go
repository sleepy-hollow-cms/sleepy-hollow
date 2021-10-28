package usecase

import (
	"context"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/port"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/util/log"
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

func (u *User) Update(user domain.User) (*domain.User, error) {
	_, err := u.UserPort.FindById(context.TODO(), user.Id)
	if err != nil {
		log.Logger.Warn(err.Error())
		return nil, err
	}

	updated := domain.User{
		Id:   user.Id,
		Name: user.Name,
	}

	result, err := u.UserPort.Update(context.TODO(), updated)
	if err != nil {
		log.Logger.Warn(err.Error())
		return nil, err
	}

	return result, nil
}
