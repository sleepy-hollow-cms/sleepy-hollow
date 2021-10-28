package gateway

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/usecase"
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

	t.Run("UserをID指定で取得する", func(t *testing.T) {
		mockUserDriver := new(MockUserDriver)
		mockUserDriver.On("FindById", "id").Return(&model.User{
			Id:   "id",
			Name: "name",
		}, nil)

		target.UserDriver = mockUserDriver

		id := domain.UserId("id")
		actual, err := target.FindById(context.TODO(), id)

		assert.Nil(t, err)
		assert.Equal(t, &domain.User{
			Id:   domain.UserId("id"),
			Name: domain.UserName("name"),
		}, actual)
	})

	t.Run("UserをID指定で更新できる", func(t *testing.T) {
		mockUserDriver := new(MockUserDriver)

		user := model.User{
			Id:   "userid",
			Name: "username",
		}

		mockUserDriver.On("Update", mock.AnythingOfType("model.User")).Return(&user, nil)

		target.UserDriver = mockUserDriver

		updatedUser := domain.User{
			Id:   domain.UserId("userid"),
			Name: domain.UserName("username"),
		}

		actual, err := target.Update(context.TODO(), updatedUser)

		expected := &domain.User{
			Id:   domain.UserId("userid"),
			Name: domain.UserName("username"),
		}

		mockUserDriver.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("IDが存在しない時Userが更新できない", func(t *testing.T) {
		mockUserDriver := new(MockUserDriver)

		updatedUser := domain.User{
			Id:   domain.UserId("userid"),
			Name: domain.UserName("username"),
		}

		mockUserDriver.On("Update", mock.AnythingOfType("model.User")).Return(nil, driver.NewUserCannotUpdateError())
		target.UserDriver = mockUserDriver

		_, err := target.Update(context.TODO(), updatedUser)

		expected := usecase.NewUserUpdateFailedError("User Update failed")

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &expected))
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

func (_m *MockUserDriver) FindById(userId string) (*model.User, error) {
	ret := _m.Called(userId)
	user := ret.Get(0)
	if user != nil {
		return user.(*model.User), ret.Error(1)
	}
	return nil, ret.Error(1)
}

func (_m *MockUserDriver) Update(user model.User) (*model.User, error) {
	ret := _m.Called(user)
	userModel := ret.Get(0)
	if userModel != nil {
		return userModel.(*model.User), ret.Error(1)
	}
	return nil, ret.Error(1)
}
