package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/usecase"
)

func TestEntry(t *testing.T) {
	var target = usecase.Entry{}

	t.Run("Entryを登録することができる", func(t *testing.T) {
		mockEntryPort := new(MockEntryPort)
		mockContentModelPort := new(MockContentModelPort)
		modelID := domain.ContentModelID("modelId")
		entry := domain.Entry{
			ContentModelID: modelID,
			ID:             domain.EntryId("id"),
		}

		entryItems := []domain.EntryItem{
			{
				Value: "タイトル1",
			},
		}

		inputEntryItems := []domain.EntryItem{
			{
				Value: "タイトル1",
			},
		}

		inputEntry := domain.Entry{
			ContentModelID: domain.ContentModelID("modelId"),
			Items:          inputEntryItems,
		}

		model := domain.ContentModel{
			Fields: domain.Fields{
				{
					Name:     domain.Name("name"),
					Type:     domain.Text,
					Required: domain.Required(true),
				},
			},
		}

		mockEntryPort.On("Create", inputEntry).Return(entry)
		mockEntryPort.On("CreateItems", domain.EntryId("id"), inputEntryItems).Return(entryItems, nil)
		mockContentModelPort.On("FindByID", modelID).Return(model, nil)

		target.EntryPort = mockEntryPort
		target.ContentModelPort = mockContentModelPort

		expected := domain.Entry{
			ID:             domain.EntryId("id"),
			ContentModelID: domain.ContentModelID("modelId"),
			Items:          entryItems,
		}
		actual, err := target.Register(inputEntry)

		mockEntryPort.AssertExpectations(t)
		mockContentModelPort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("ContentModelIDが存在しないものであった場合はErrorを返す", func(t *testing.T) {

		mockEntryPort := new(MockEntryPort)
		mockContentModelPort := new(MockContentModelPort)

		id := domain.ContentModelID("id")

		inputEntryItems := []domain.EntryItem{
			{
				Value: domain.TextValue{
					Value: "タイトル1",
				},
			},
			{
				Value: domain.TextValue{
					Value: "タイトル1",
				},
			},
		}

		entry := domain.Entry{
			ContentModelID: id,
			Items:          inputEntryItems,
		}

		contentModelNotFound := usecase.NewContentModelNotFoundError("content model not found")
		mockContentModelPort.On("FindByID", id).Return(domain.ContentModel{}, contentModelNotFound)

		target.EntryPort = mockEntryPort
		target.ContentModelPort = mockContentModelPort

		_, err := target.Register(entry)

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &contentModelNotFound))
		mockEntryPort.AssertNotCalled(t, "FindByID")
	})

	t.Run("EntryがContentModelの形にあわない場合はErrorを返す", func(t *testing.T) {
		mockEntryPort := new(MockEntryPort)
		mockContentModelPort := new(MockContentModelPort)
		modelID := domain.ContentModelID("modelId")
		entry := domain.Entry{
			ContentModelID: modelID,
			ID:             domain.EntryId("id"),
		}

		entryItems := []domain.EntryItem{
			{
				Value: domain.TextValue{
					Value: "タイトル1",
				},
			},
		}

		inputEntryItems := []domain.EntryItem{
			{
				Value: domain.TextValue{
					Value: "タイトル1",
				},
			},
			{
				Value: domain.MultipleTextValue{
					Value: []string{
						"text1",
						"text2",
					},
				},
			},
		}

		inputEntry := domain.Entry{
			ContentModelID: domain.ContentModelID("modelId"),
			Items:          inputEntryItems,
		}

		model := domain.ContentModel{
			ID: modelID,
			Fields: domain.Fields{
				{
					Name:     domain.Name("name"),
					Type:     domain.Text,
					Required: domain.Required(true),
				},
			},
		}

		mockEntryPort.On("Create", inputEntry).Return(entry)
		mockEntryPort.On("CreateItems", domain.EntryId("id"), inputEntryItems).Return(entryItems, nil)
		mockContentModelPort.On("FindByID", modelID).Return(model, nil)

		target.EntryPort = mockEntryPort
		target.ContentModelPort = mockContentModelPort

		_, err := target.Register(inputEntry)

		entryValidationError := domain.NewEntryValidationError("test reason")

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &entryValidationError))
		mockEntryPort.AssertNotCalled(t, "FindByID")
	})

	t.Run("不明なエラーが返された場合はそのままに返す", func(t *testing.T) {

		mockEntryPort := new(MockEntryPort)
		mockContentModelPort := new(MockContentModelPort)

		id := domain.ContentModelID("id")
		entry := domain.Entry{
			ContentModelID: id,
			Items: []domain.EntryItem{
				{
					Value: domain.TextValue{
						Value: "タイトル1",
					},
				},
				{
					Value: domain.TextValue{
						Value: "タイトル1",
					},
				},
			},
		}

		someError := errors.New("some error")
		mockContentModelPort.On("FindByID", id).Return(domain.ContentModel{}, someError)

		target.EntryPort = mockEntryPort
		target.ContentModelPort = mockContentModelPort

		_, err := target.Register(entry)

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &someError))
		mockEntryPort.AssertNotCalled(t, "FindByID")
	})

	t.Run("EntryをID指定で取得することができる", func(t *testing.T) {
		id := domain.EntryId("id")

		mockEntryPort := new(MockEntryPort)

		expected := domain.Entry{
			ID:             domain.EntryId("entryId"),
			ContentModelID: domain.ContentModelID("contentModelId"),
			Items: []domain.EntryItem{
				{
					Value: domain.TextValue{Value: "タイトル"},
				},
			},
		}

		mockEntryPort.On("FindById", id).Return(expected, nil)

		target.EntryPort = mockEntryPort

		actual, err := target.FindByID(id)

		mockEntryPort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("EntryをID指定で取得できない場合EntryNotFoundErrorを返す", func(t *testing.T) {
		id := domain.EntryId("id")

		mockEntryPort := new(MockEntryPort)

		mockEntryPort.On("FindById", id).Return(domain.Entry{}, errors.New("not found"))

		target.EntryPort = mockEntryPort

		_, err := target.FindByID(id)

		notFoundError := usecase.NewEntryNotFoundError("error")

		mockEntryPort.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &notFoundError))
	})
}

type MockEntryPort struct {
	mock.Mock
}

func (_m *MockEntryPort) CreateItems(ctx context.Context, id domain.EntryId, entry []domain.EntryItem) ([]domain.EntryItem, error) {
	ret := _m.Called(id, entry)
	return ret.Get(0).([]domain.EntryItem), nil
}

func (_m *MockEntryPort) Create(ctx context.Context, entry domain.Entry) (domain.Entry, error) {
	ret := _m.Called(entry)
	return ret.Get(0).(domain.Entry), nil
}

func (_m *MockEntryPort) FindById(ctx context.Context, id domain.EntryId) (domain.Entry, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.Entry), ret.Error(1)
}

func (_m *MockEntryPort) Find(ctx context.Context) (domain.Entries, error) {
	ret := _m.Called()
	return ret.Get(0).(domain.Entries), ret.Error(1)
}
