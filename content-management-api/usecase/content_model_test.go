package usecase_test

import (
	"content-management-api/domain"
	"content-management-api/support"
	"content-management-api/usecase"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestContentModel(t *testing.T) {

	var target = usecase.ContentModelUseCase{}

	t.Run("ContentModelをIDを使って取得することができる", func(t *testing.T) {
		id := domain.ContentModelID("id")

		// Mock setting
		mockContentModelPort := new(MockContentModelPort)
		modelID := domain.ContentModelID("id")
		model := domain.ContentModel{
			ID: id,
		}
		mockContentModelPort.On("FindByID", modelID).Return(model, nil)
		target.ContentModelPort = mockContentModelPort

		expected := domain.ContentModel{
			ID: id,
		}

		actual, err := target.FindContentModelByID(id)

		mockContentModelPort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("存在しないContentModelのIDの場合はContentModelNotFoundErrorを返す", func(t *testing.T) {
		id := domain.ContentModelID("id")

		// Mock setting
		mockContentModelPort := new(MockContentModelPort)
		modelID := domain.ContentModelID("id")
		model := domain.ContentModel{
			ID: id,
		}

		contentModelNotFoundError := usecase.ContentModelNotFoundError{Reason: "test"}

		mockContentModelPort.On("FindByID", modelID).Return(model, &contentModelNotFoundError)
		target.ContentModelPort = mockContentModelPort

		_, err := target.FindContentModelByID(id)

		mockContentModelPort.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &contentModelNotFoundError))
	})

	t.Run("ContentModelを登録することができる", func(t *testing.T) {
		// Mock setting
		mockContentModelPort := new(MockContentModelPort)
		contentModel := domain.ContentModel{
			ID: domain.ContentModelID("id"),
		}
		mockContentModelPort.On("Save", contentModel).Return(contentModel, nil)
		target.ContentModelPort = mockContentModelPort

		expected := domain.ContentModel{
			ID: domain.ContentModelID("id"),
		}

		actual, err := target.CreateContentModel(contentModel)

		mockContentModelPort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("ContentModelを登録できない場合はContentModelSaveFailedErrorを返す", func(t *testing.T) {
		// Mock setting
		mockContentModelPort := new(MockContentModelPort)
		contentModel := domain.ContentModel{
			ID: domain.ContentModelID("id"),
		}

		saveFailError := usecase.ContentModelSaveFailError{Reason: "test"}
		mockContentModelPort.On("Save", contentModel).Return(domain.ContentModel{}, &saveFailError)
		target.ContentModelPort = mockContentModelPort

		_, err := target.CreateContentModel(contentModel)

		mockContentModelPort.AssertExpectations(t)
		assert.True(t, errors.As(err, &saveFailError))
	})

	t.Run("Spaceに紐づくContentModelを全て取得することができる", func(t *testing.T) {
		t.Fatalf(support.TODO("Test code needs to be implemented.").Error())
	})
}

type MockContentModelPort struct {
	mock.Mock
}

func (_m *MockContentModelPort) FindByID(ctx context.Context, id domain.ContentModelID) (domain.ContentModel, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.ContentModel), ret.Error(1)
}

func (_m *MockContentModelPort) FindBySpaceID(ctx context.Context, id domain.SpaceID) (domain.ContentModels, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.ContentModels), ret.Error(1)
}

func (_m *MockContentModelPort) Save(ctx context.Context, contentModel domain.ContentModel) error {
	ret := _m.Called(contentModel)
	return ret.Error(1)
}
