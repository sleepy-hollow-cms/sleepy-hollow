package gateway

import (
	"context"
	"errors"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/usecase"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/util/log"
)

type User struct {
	UserDriver driver.UserDriver
}

func NewUser(driver driver.UserDriver) *User {
	return &User{
		UserDriver: driver,
	}
}

func (u *User) Register(ctx context.Context, user domain.User) (domain.User, error) {
	input := model.User{
		Name: user.Name.String(),
	}

	registeredUser, err := u.UserDriver.Register(input)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		Id:   domain.UserId(registeredUser.Id),
		Name: domain.UserName(registeredUser.Name),
	}, nil
}

func (u *User) DeleteById(ctx context.Context, userId domain.UserId) error {
	_, err := u.UserDriver.DeleteById(userId.String())

	if err != nil {
		switch err.(type) {
		case driver.UserNotFoundError:
			return err
		default:
			return err
		}
	}
	return nil
}

func (u *User) FindById(ctx context.Context, userId domain.UserId) (*domain.User, error) {
	user, err := u.UserDriver.FindById(userId.String())

	if err != nil {
		log.Logger.Warn(err.Error())
		switch {
		case errors.As(err, &driver.UserNotFoundError{}):
			return nil, usecase.NewUserNotFoundError("User Not Found")
		default:
			return nil, err
		}
	}

	return &domain.User{
		Id:   domain.UserId(user.Id),
		Name: domain.UserName(user.Name),
	}, nil
}

func (u *User) Update(ctx context.Context, updatedUser domain.User) (*domain.User, error) {
	updated, err := u.UserDriver.Update(
		model.User{
			Id:   updatedUser.Id.String(),
			Name: updatedUser.Name.String(),
		})

	if err != nil {
		log.Logger.Warn(err.Error())
		switch {
		case errors.As(err, &driver.UserCannotUpdateError{}):
			return nil, usecase.NewUserUpdateFailedError("Content User Update conflicted")
		default:
			return nil, err
		}
	}

	return &domain.User{
		Id:   domain.UserId(updated.Id),
		Name: domain.UserName(updated.Name),
	}, nil
}
