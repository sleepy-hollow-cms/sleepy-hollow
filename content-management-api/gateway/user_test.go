package gateway

import (
	"context"
	"testing"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUser(t *testing.T) {
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

	t.Run("UserをID指定で削除する", func(t *testing.T) {
		mockUserDriver := new(MockUserDriver)
		mockUserDriver.On("DeleteById", "id").Return(int64(1), nil)

		target.UserDriver = mockUserDriver

		id := domain.UserId("id")
		err := target.DeleteById(context.TODO(), id)

		assert.Nil(t, err)
	})
}

type MockUserDriver struct {
	mock.Mock
}

func (_m *MockUserDriver) Register(user model.User) (model.User, error) {
	ret := _m.Called(user)
	return ret.Get(0).(model.User), ret.Error(1)
}

func (_m *MockUserDriver) DeleteById(userId string) (int64, error) {
	ret := _m.Called(userId)
	return ret.Get(0).(int64), ret.Error(1)
}
