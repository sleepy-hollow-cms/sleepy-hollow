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

	t.Run("Spaceを全て取得することができる", func(t *testing.T) {
		// Mock setting
		mockSpacePort := new(MockSpacePort)
		spaces := domain.Spaces{
			{ID: domain.SpaceID("id1"), Name: domain.Name("name1")},
			{ID: domain.SpaceID("id2"), Name: domain.Name("name2")},
		}
		mockSpacePort.On("Find").Return(spaces, nil)
		target.SpacePort = mockSpacePort

		expected := domain.Spaces{
			{ID: domain.SpaceID("id1"), Name: domain.Name("name1")},
			{ID: domain.SpaceID("id2"), Name: domain.Name("name2")},
		}

		actual, err := target.Find()

		mockSpacePort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Spaceを取得したときSpaceが一つも取れなかった場合は空配列", func(t *testing.T) {
		// Mock setting
		mockSpacePort := new(MockSpacePort)
		spaces := make(domain.Spaces, 0)
		mockSpacePort.On("Find").Return(spaces, nil)
		target.SpacePort = mockSpacePort

		expected := make(domain.Spaces, 0)

		actual, err := target.Find()

		mockSpacePort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Spaceを取得したときエラーが発生した場合はエラーをそのまま返す", func(t *testing.T) {
		// Mock setting
		mockSpacePort := new(MockSpacePort)
		somethingError := errors.New("something error")
		mockSpacePort.On("Find").Return(nil, somethingError)
		target.SpacePort = mockSpacePort

		result, err := target.Find()

		mockSpacePort.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, somethingError, err)
	})

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

func (_m *MockSpacePort) Find(ctx context.Context) (domain.Spaces, error) {
	ret := _m.Called()

	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}

	return ret.Get(0).(domain.Spaces), ret.Error(1)
}

func (_m *MockSpacePort) FindByID(ctx context.Context, id domain.SpaceID) (domain.Space, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.Space), ret.Error(1)
}

func (_m *MockSpacePort) Register(ctx context.Context, space domain.Space) (domain.Space, error) {
	ret := _m.Called(space)
	return ret.Get(0).(domain.Space), ret.Error(1)
}

func (_m *MockSpacePort) Update(ctx context.Context, space domain.Space) (domain.Space, error) {
	ret := _m.Called(space)
	return ret.Get(0).(domain.Space), ret.Error(1)
}
