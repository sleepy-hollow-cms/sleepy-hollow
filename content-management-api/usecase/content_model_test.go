package usecase_test

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/usecase"
	"content-management-api/usecase/write"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestContentModel(t *testing.T) {

	var target = usecase.ContentModel{}

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

		actual, err := target.FindByID(id)

		mockContentModelPort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("存在しないContentModelのIDを取得した場合はContentModelNotFoundErrorを返す", func(t *testing.T) {
		id := domain.ContentModelID("id")

		// Mock setting
		mockContentModelPort := new(MockContentModelPort)
		modelID := domain.ContentModelID("id")
		model := domain.ContentModel{
			ID: id,
		}

		contentModelNotFoundError := usecase.NewContentModelNotFoundError("test")

		mockContentModelPort.On("FindByID", modelID).Return(model, contentModelNotFoundError)
		target.ContentModelPort = mockContentModelPort

		_, err := target.FindByID(id)

		mockContentModelPort.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &contentModelNotFoundError))
	})

	t.Run("ContentModelをIDを使って削除することができる", func(t *testing.T) {
		id := domain.ContentModelID("id")

		// Mock setting
		mockContentModelPort := new(MockContentModelPort)
		modelID := domain.ContentModelID("id")

		mockContentModelPort.On("DeleteByID", modelID).Return(nil)
		target.ContentModelPort = mockContentModelPort

		err := target.DeleteContentModelByID(id)

		mockContentModelPort.AssertExpectations(t)
		assert.Nil(t, err)
	})

	t.Run("存在しないContentModelのIDを削除した場合はContentModelNotFoundErrorを返す", func(t *testing.T) {
		id := domain.ContentModelID("id")

		// Mock setting
		mockContentModelPort := new(MockContentModelPort)
		modelID := domain.ContentModelID("id")

		contentModelNotFoundError := usecase.NewContentModelNotFoundError("test")

		mockContentModelPort.On("DeleteByID", modelID).Return(contentModelNotFoundError)
		target.ContentModelPort = mockContentModelPort

		err := target.DeleteContentModelByID(id)

		mockContentModelPort.AssertExpectations(t)
		assert.True(t, errors.As(err, &contentModelNotFoundError))
	})

	t.Run("Spaceに紐づくContentModelを全て取得することができる", func(t *testing.T) {
		spaceID := domain.SpaceID("spaceID")
		// Mock setting
		mockContentModelPort := new(MockContentModelPort)
		contentModels := domain.NewContentModels([]domain.ContentModel{
			{ID: domain.ContentModelID("id")},
		})
		mockContentModelPort.On("FindBySpaceID", spaceID).Return(contentModels, nil)
		target.ContentModelPort = mockContentModelPort

		expected := domain.NewContentModels([]domain.ContentModel{
			{ID: domain.ContentModelID("id")},
		})

		actual, err := target.FindContentModelBySpaceID(spaceID)

		mockContentModelPort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Spaceに紐づくContentModelの取得に失敗した場合はErrorを返す", func(t *testing.T) {
		spaceID := domain.SpaceID("spaceID")
		// Mock setting
		mockContentModelPort := new(MockContentModelPort)
		contentModels := domain.NewContentModels([]domain.ContentModel{
			{ID: domain.ContentModelID("id")},
		})
		mockContentModelPort.On("FindBySpaceID", spaceID).Return(contentModels, errors.New("error"))
		target.ContentModelPort = mockContentModelPort

		_, err := target.FindContentModelBySpaceID(spaceID)

		mockContentModelPort.AssertExpectations(t)
		assert.NotNil(t, err)
	})

	t.Run("Spaceに紐づくContentModelsが一件もない場合はサイズが0のコレクションを返す", func(t *testing.T) {
		spaceID := domain.SpaceID("spaceID")
		// Mock setting
		mockContentModelPort := new(MockContentModelPort)
		contentModels := domain.NewContentModels([]domain.ContentModel{})
		mockContentModelPort.On("FindBySpaceID", spaceID).Return(contentModels, nil)
		target.ContentModelPort = mockContentModelPort

		actual, err := target.FindContentModelBySpaceID(spaceID)

		mockContentModelPort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.True(t, len(actual) == 0)
	})

	t.Run("ContentModelを登録することができる", func(t *testing.T) {
		contentModel := write.ContentModel{
			Name: domain.Name("name"),
			Fields: []field.Field{
				{
					Type:     field.Text,
					Required: field.Required(true),
				},
			},
		}

		retContentModel := domain.ContentModel{
			ID:   domain.ContentModelID("id"),
			Name: domain.Name("name"),
			Fields: []field.Field{
				{
					Type:     field.Text,
					Required: field.Required(true),
				},
			},
		}

		// Mock setting
		mockContentModelPort := new(MockContentModelPort)
		mockContentModelPort.On("Create", contentModel).Return(retContentModel, nil)
		target.ContentModelPort = mockContentModelPort

		actual, err := target.Create(contentModel)

		expected := domain.ContentModel{
			ID:   domain.ContentModelID("id"),
			Name: domain.Name("name"),
			Fields: []field.Field{
				{
					Type:     field.Text,
					Required: field.Required(true),
				},
			},
		}

		mockContentModelPort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("ContentModelを登録時に失敗した場合はContentModelCreateFailedErrorを返す", func(t *testing.T) {
		contentModel := write.ContentModel{
			Fields: []field.Field{
				{Type: field.Text},
			},
		}

		retContentModel := domain.ContentModel{
			ID:     domain.ContentModelID("id"),
			Fields: nil,
		}

		contentModelCreateFailedError := usecase.NewContentModelCreateFailedError("test")

		// Mock setting
		mockContentModelPort := new(MockContentModelPort)
		mockContentModelPort.On("Create", contentModel).Return(retContentModel, &contentModelCreateFailedError)
		target.ContentModelPort = mockContentModelPort

		_, err := target.Create(contentModel)

		mockContentModelPort.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &contentModelCreateFailedError))
	})

}

type MockContentModelPort struct {
	mock.Mock
}

func (_m *MockContentModelPort) FindByID(ctx context.Context, id domain.ContentModelID) (domain.ContentModel, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.ContentModel), ret.Error(1)
}

func (_m *MockContentModelPort) DeleteByID(ctx context.Context, id domain.ContentModelID) error {
	ret := _m.Called(id)
	return ret.Error(0)
}

func (_m *MockContentModelPort) FindBySpaceID(ctx context.Context, id domain.SpaceID) (domain.ContentModels, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.ContentModels), ret.Error(1)
}

func (_m *MockContentModelPort) Save(ctx context.Context, contentModel domain.ContentModel) error {
	ret := _m.Called(contentModel)
	return ret.Error(1)
}

func (_m *MockContentModelPort) Create(ctx context.Context, contentModel write.ContentModel) (domain.ContentModel, error) {
	ret := _m.Called(contentModel)
	return ret.Get(0).(domain.ContentModel), ret.Error(1)
}
