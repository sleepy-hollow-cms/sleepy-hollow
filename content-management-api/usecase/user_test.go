package usecase_test

import (
	"context"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
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
}

type MockUserPort struct {
	mock.Mock
}

func (_m MockUserPort) Register(ctx context.Context, user domain.User) (domain.User, error) {
	ret := _m.Called(user)
	return ret.Get(0).(domain.User), ret.Error(1)
}
