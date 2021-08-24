package usecase_test

import (
	"content-management-api/domain"
	"content-management-api/usecase"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSpace(t *testing.T) {

	var target = usecase.Space{}

	t.Run("SpaceをIDで取得することができる", func(t *testing.T) {
		id := domain.SpaceID("id")

		// Mock setting
		mockSpacePort := new(MockSpacePort)
		spaceID := domain.SpaceID("id")
		model := domain.Space{
			ID: id,
		}
		mockSpacePort.On("FindByID", spaceID).Return(model, nil)
		target.SpacePort = mockSpacePort

		expected := domain.Space{
			ID: id,
		}

		actual, err := target.FindByID(id)

		mockSpacePort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("存在しないSpaceIDを指定して取得しようした場合はSpaceNotFoundErrorを返す", func(t *testing.T) {
		id := domain.SpaceID("id")

		// Mock setting
		mockSpacePort := new(MockSpacePort)
		modelID := domain.SpaceID("id")

		spaceNotFoundError := usecase.NewSpaceNotFoundError("test")

		mockSpacePort.On("FindByID", modelID).Return(domain.Space{}, &spaceNotFoundError)
		target.SpacePort = mockSpacePort

		_, err := target.FindByID(id)

		mockSpacePort.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &spaceNotFoundError))
	})

	t.Run("Spaceの登録をすることができる", func(t *testing.T) {
		mockSpacePort := new(MockSpacePort)
		space := domain.Space{
			Name: domain.Name("name"),
		}

		mockSpacePort.On("Register", space).Return(domain.Space{
			ID:   domain.SpaceID("spaceid"),
			Name: domain.Name("name"),
		}, nil)
		target.SpacePort = mockSpacePort

		actual, err := target.Register(domain.Space{
			Name: "name",
		})

		expected := domain.Space{
			ID:   domain.SpaceID("spaceid"),
			Name: domain.Name("name"),
		}

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

}

type MockSpacePort struct {
	mock.Mock
}

func (_m *MockSpacePort) FindByID(ctx context.Context, id domain.SpaceID) (domain.Space, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.Space), ret.Error(1)
}

func (_m *MockSpacePort) Register(ctx context.Context, space domain.Space) (domain.Space, error) {
	ret := _m.Called(space)
	return ret.Get(0).(domain.Space), ret.Error(1)
}
