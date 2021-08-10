package usecase_test

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/usecase"
	"content-management-api/usecase/read"
	"content-management-api/usecase/write"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
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

		entryItems := read.EntryItem{
			Items: []read.Item{
				{
					Type:      field.Text,
					FieldName: "fieldName1",
					Value: field.TextValue{
						Value: "タイトル1",
					},
				},
			},
		}

		inputEntry := write.Entry{
			ContentModelID: domain.ContentModelID("modelId"),
		}

		inputEntryItems := []write.EntryItem{
			{
				Type:      field.Text,
				FieldName: "fieldName1",
				Value: field.TextValue{
					Value: "タイトル1",
				},
			},
		}

		model := domain.ContentModel{}

		mockEntryPort.On("Create", inputEntry).Return(entry)
		mockEntryPort.On("CreateItems", domain.EntryId("id"), inputEntryItems).Return(entryItems, nil)
		mockContentModelPort.On("FindByID", modelID).Return(model, nil)

		target.EntryPort = mockEntryPort
		target.ContentModelPort = mockContentModelPort

		expected := read.Entry{
			ID:             domain.EntryId("id"),
			ContentModelID: domain.ContentModelID("modelId"),
			EntryItems:     entryItems,
		}
		actual, err := target.Register(inputEntry, inputEntryItems)

		mockEntryPort.AssertExpectations(t)
		mockContentModelPort.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("ContentModelIDが存在しないものであった場合はErrorを返す", func(t *testing.T) {

		mockEntryPort := new(MockEntryPort)
		mockContentModelPort := new(MockContentModelPort)

		id := domain.ContentModelID("id")
		entry := write.Entry{
			ContentModelID: id,
		}

		inputEntryItems := []write.EntryItem{
			{
				FieldName: "fieldName1",
				Value: field.TextValue{
					Value: "タイトル1",
				},
			},
			{
				FieldName: "fieldName2",
				Value: field.TextValue{
					Value: "タイトル1",
				},
			},
		}

		contentModelNotFound := usecase.NewContentModelNotFoundError("content model not found")
		mockContentModelPort.On("FindByID", id).Return(domain.ContentModel{}, contentModelNotFound)

		target.EntryPort = mockEntryPort
		target.ContentModelPort = mockContentModelPort

		_, err := target.Register(entry, inputEntryItems)

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &contentModelNotFound))
		mockEntryPort.AssertNotCalled(t, "FindByID")
	})

	t.Run("不明なエラーが返された場合はそのままに返す", func(t *testing.T) {

		mockEntryPort := new(MockEntryPort)
		mockContentModelPort := new(MockContentModelPort)

		id := domain.ContentModelID("id")
		entry := write.Entry{
			ContentModelID: id,
		}

		inputEntryItems := []write.EntryItem{
			{
				FieldName: "fieldName1",
				Value: field.TextValue{
					Value: "タイトル1",
				},
			},
			{
				FieldName: "fieldName2",
				Value: field.TextValue{
					Value: "タイトル1",
				},
			},
		}

		someError := errors.New("some error")
		mockContentModelPort.On("FindByID", id).Return(domain.ContentModel{}, someError)

		target.EntryPort = mockEntryPort
		target.ContentModelPort = mockContentModelPort

		_, err := target.Register(entry, inputEntryItems)

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &someError))
		mockEntryPort.AssertNotCalled(t, "FindByID")
	})
}

type MockEntryPort struct {
	mock.Mock
}

func (_m *MockEntryPort) CreateItems(ctx context.Context, id domain.EntryId, entry []write.EntryItem) (read.EntryItem, error) {
	ret := _m.Called(id, entry)
	return ret.Get(0).(read.EntryItem), nil
}

func (_m *MockEntryPort) Create(ctx context.Context, entry write.Entry) (domain.Entry, error) {
	ret := _m.Called(entry)
	return ret.Get(0).(domain.Entry), nil
}
