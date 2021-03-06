package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		modelID := domain.ContentModelID("id")
		mockContentModelPort := new(MockContentModelPort)
		mockContentModelPort.On("DeleteByID", modelID).Return(nil)
		target.ContentModelPort = mockContentModelPort

		mockEntryPort := new(MockEntryPort)
		mockEntryPort.On("Find").Return(domain.Entries{}, nil)
		target.EntryPort = mockEntryPort

		err := target.DeleteContentModelByID(id)

		mockContentModelPort.AssertExpectations(t)
		assert.Nil(t, err)
	})

	t.Run("ContentModelをIDを使ってEntryが存在する場合はReferenceByEntryErrorを返す", func(t *testing.T) {
		id := domain.ContentModelID("id")

		// Mock setting
		modelID := domain.ContentModelID("id")
		mockContentModelPort := new(MockContentModelPort)
		mockContentModelPort.On("DeleteByID", modelID).Return(nil)
		target.ContentModelPort = mockContentModelPort

		mockEntryPort := new(MockEntryPort)
		mockEntryPort.On("Find").Return(domain.Entries{
			{
				ContentModelID: domain.ContentModelID("id"),
			},
		}, nil)
		target.EntryPort = mockEntryPort

		err := target.DeleteContentModelByID(id)

		mockContentModelPort.AssertNotCalled(t, "DeleteByID")
		mockEntryPort.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &usecase.ReferenceByEntryError{}))
	})

	t.Run("存在しないContentModelのIDを削除した場合はContentModelNotFoundErrorを返す", func(t *testing.T) {
		id := domain.ContentModelID("id")

		mockContentModelPort := new(MockContentModelPort)
		modelID := domain.ContentModelID("id")

		contentModelNotFoundError := usecase.NewContentModelNotFoundError("test")

		mockContentModelPort.On("DeleteByID", modelID).Return(contentModelNotFoundError)
		target.ContentModelPort = mockContentModelPort

		mockEntryPort := new(MockEntryPort)
		mockEntryPort.On("Find").Return(domain.Entries{}, nil)
		target.EntryPort = mockEntryPort

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
		assert.True(t, len(actual.Models) == 0)
	})

	t.Run("UpdatedAtが一致している時ContentModelを更新することができる", func(t *testing.T) {
		rowCreatedTime := time.Now()
		createdAt := domain.CreatedAt(rowCreatedTime)
		updatedAt := domain.UpdatedAt(rowCreatedTime)

		id := domain.ContentModelID("id")

		foundContentModel := domain.ContentModel{
			ID:        id,
			Name:      domain.Name("name"),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Fields: []domain.Field{
				{
					Type:     domain.Text,
					Name:     domain.Name("fieldName"),
					Required: domain.Required(true),
				},
			},
		}

		updatedContentModel := domain.ContentModel{
			ID:        id,
			Name:      domain.Name("updated_name"),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Fields: []domain.Field{
				{
					Type:     domain.Number,
					Name:     domain.Name("number"),
					Required: domain.Required(false),
				},
			},
		}

		result := domain.ContentModel{
			ID:        id,
			Name:      domain.Name("updated_name"),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Fields: []domain.Field{
				{
					Type:     domain.Number,
					Name:     domain.Name("number"),
					Required: domain.Required(false),
				},
			},
		}

		mockContentModelPort := new(MockContentModelPort)

		mockContentModelPort.On("FindByID", id).Return(foundContentModel, nil)
		mockContentModelPort.On("Update", updatedContentModel).Return(result, nil)
		target.ContentModelPort = mockContentModelPort

		actual, err := target.Update(updatedContentModel)

		expected := domain.ContentModel{
			ID:        id,
			Name:      domain.Name("updated_name"),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Fields: []domain.Field{
				{
					Type:     domain.Number,
					Name:     domain.Name("number"),
					Required: domain.Required(false),
				},
			},
		}

		mockContentModelPort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("指定したIDが存在しない時ContentModelを更新することができずErrorを返す", func(t *testing.T) {
		rowCreatedTime := time.Now()
		createdAt := domain.CreatedAt(rowCreatedTime)
		updatedAt := domain.UpdatedAt(rowCreatedTime)

		id := domain.ContentModelID("id")

		updatedContentModel := domain.ContentModel{
			ID:        id,
			Name:      domain.Name("updated_name"),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Fields: []domain.Field{
				{
					Type:     domain.Number,
					Name:     domain.Name("number"),
					Required: domain.Required(false),
				},
			},
		}

		mockContentModelPort := new(MockContentModelPort)

		mockContentModelPort.On("FindByID", id).Return(domain.ContentModel{}, usecase.NewContentModelNotFoundError("Content Model Not Found"))
		target.ContentModelPort = mockContentModelPort

		_, err := target.Update(updatedContentModel)

		expected := usecase.NewContentModelNotFoundError("Content Model Not Found")

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &expected))
	})

	t.Run("UpdatedAtが一致していない時ContentModelを更新することができずErrorを返す", func(t *testing.T) {
		rowCreatedTime := time.Now()
		createdAt := domain.CreatedAt(rowCreatedTime)
		currentUpdateAt := domain.UpdatedAt(rowCreatedTime)

		updatedAt := domain.UpdatedAt(time.Now())

		id := domain.ContentModelID("id")

		foundContentModel := domain.ContentModel{
			ID:        id,
			Name:      domain.Name("name"),
			CreatedAt: createdAt,
			UpdatedAt: currentUpdateAt,
			Fields: []domain.Field{
				{
					Type:     domain.Text,
					Name:     domain.Name("fieldName"),
					Required: domain.Required(true),
				},
			},
		}

		updatedContentModel := domain.ContentModel{
			ID:        id,
			Name:      domain.Name("updated_name"),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Fields: []domain.Field{
				{
					Type:     domain.Number,
					Name:     domain.Name("number"),
					Required: domain.Required(false),
				},
			},
		}

		result := domain.ContentModel{}

		mockContentModelPort := new(MockContentModelPort)

		mockContentModelPort.On("FindByID", id).Return(foundContentModel, nil)
		mockContentModelPort.On("Update", updatedContentModel).Return(result, usecase.NewContentModelUpdateFailedError("Content Model Update conflicted"))
		target.ContentModelPort = mockContentModelPort

		_, err := target.Update(updatedContentModel)

		expected := usecase.NewContentModelUpdateFailedError("Content Model Update conflicted")

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &expected))
	})

	t.Run("ContentModelを登録することができる", func(t *testing.T) {

		createdAt := domain.CreatedAt(time.Now())

		contentModel := domain.ContentModel{
			Name:      domain.Name("name"),
			CreatedAt: createdAt,
			Fields: []domain.Field{
				{
					Type:     domain.Text,
					Name:     domain.Name("fieldName"),
					Required: domain.Required(true),
				},
			},
		}

		retContentModel := domain.ContentModel{
			ID:        domain.ContentModelID("id"),
			Name:      domain.Name("name"),
			CreatedAt: createdAt,
			Fields: []domain.Field{
				{
					Type:     domain.Text,
					Name:     domain.Name("fieldName"),
					Required: domain.Required(true),
				},
			},
		}

		// Mock setting
		mockContentModelPort := new(MockContentModelPort)
		mockContentModelPort.On("Create", contentModel).Return(retContentModel, nil)
		target.ContentModelPort = mockContentModelPort

		actual, err := target.Create(contentModel)

		expected := domain.ContentModel{
			ID:        domain.ContentModelID("id"),
			Name:      domain.Name("name"),
			CreatedAt: createdAt,
			Fields: []domain.Field{
				{
					Type:     domain.Text,
					Name:     domain.Name("fieldName"),
					Required: domain.Required(true),
				},
			},
		}

		mockContentModelPort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("ContentModelを登録時に失敗した場合はContentModelCreateFailedErrorを返す", func(t *testing.T) {
		contentModel := domain.ContentModel{
			Fields: []domain.Field{
				{Type: domain.Text},
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

func (_m *MockContentModelPort) Create(ctx context.Context, contentModel domain.ContentModel) (domain.ContentModel, error) {
	ret := _m.Called(contentModel)
	return ret.Get(0).(domain.ContentModel), ret.Error(1)
}

func (_m *MockContentModelPort) Update(ctx context.Context, updatedContentModel domain.ContentModel) (domain.ContentModel, error) {
	ret := _m.Called(updatedContentModel)
	return ret.Get(0).(domain.ContentModel), ret.Error(1)
}
