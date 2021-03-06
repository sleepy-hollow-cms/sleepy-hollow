package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/usecase"
)

func TestUser(t *testing.T) {

	var target = usecase.User{}

	t.Run("Userを登録することができる", func(t *testing.T) {
		input := domain.User{
			Name: domain.UserName("userName"),
		}

		expected := domain.User{
			Id:   domain.UserId("id"),
			Name: domain.UserName("userName"),
		}

		userPortMock := new(MockUserPort)

		userPortMock.On("Register", input).Return(expected, nil)
		target.UserPort = userPortMock

		actual, err := target.Register(input)

		userPortMock.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Userの登録に失敗したらerrorを返す", func(t *testing.T) {

	})

	t.Run("UserをID指定で削除することができる", func(t *testing.T) {
		id := domain.UserId("id")

		mockUserPort := new(MockUserPort)

		mockUserPort.On("DeleteById", id).Return(nil)

		target.UserPort = mockUserPort

		err := target.DeleteById(id)

		mockUserPort.AssertExpectations(t)
		assert.Nil(t, err)
	})

	t.Run("UserをID指定で更新することができる", func(t *testing.T) {
		id := domain.UserId("id")

		mockUserPort := new(MockUserPort)

		domainUser := domain.User{
			Id:   domain.UserId("id"),
			Name: domain.UserName("changedName"),
		}
		mockUserPort.On("FindById", id).Return(&domainUser, nil)
		mockUserPort.On("Update", domainUser).Return(&domain.User{
			Id:   domain.UserId("id"),
			Name: domain.UserName("changedName"),
		}, nil)

		target.UserPort = mockUserPort

		actual, err := target.Update(domain.User{
			Id:   domain.UserId("id"),
			Name: domain.UserName("changedName"),
		})

		mockUserPort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, &domain.User{
			Id:   domain.UserId("id"),
			Name: domain.UserName("changedName"),
		}, actual)
	})
}

type MockUserPort struct {
	mock.Mock
}

func (_m MockUserPort) Register(ctx context.Context, user domain.User) (domain.User, error) {
	ret := _m.Called(user)
	return ret.Get(0).(domain.User), ret.Error(1)
}

func (_m MockUserPort) DeleteById(ctx context.Context, id domain.UserId) error {
	ret := _m.Called(id)
	return ret.Error(0)
}

func (_m MockUserPort) Update(ctx context.Context, user domain.User) (*domain.User, error) {
	ret := _m.Called(user)
	return ret.Get(0).(*domain.User), ret.Error(1)
}

func (_m MockUserPort) FindById(ctx context.Context, user domain.UserId) (*domain.User, error) {
	ret := _m.Called(user)
	return ret.Get(0).(*domain.User), ret.Error(1)
}
