package gateway

import (
	"context"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
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
		case driver.EntryNotFoundError:
			return err
		default:
			return err
		}
	}
	return nil
}
