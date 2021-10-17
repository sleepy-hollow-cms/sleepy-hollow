package gateway

import (
	"context"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUser_Register(t *testing.T) {
	target := User{}

	t.Run("Userを作成する", func(t *testing.T) {

		input := domain.User{
			Name: domain.UserName("name"),
		}

		modelUser := model.User{
			Id:   "id",
			Name: "name",
		}

		inputModelUser := model.User{
			Name: "name",
		}

		mockUserDriver := new(MockUserDriver)
		mockUserDriver.On("Register", inputModelUser).Return(modelUser, nil)
		target.UserDriver = mockUserDriver

		expected := domain.User{
			Id:   "id",
			Name: "name",
		}

		actual, err := target.Register(context.TODO(), input)

		mockUserDriver.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)

	})
}

type MockUserDriver struct {
	mock.Mock
}

func (_m *MockUserDriver) Register(user model.User) (model.User, error) {
	ret := _m.Called(user)
	return ret.Get(0).(model.User), ret.Error(1)
}
