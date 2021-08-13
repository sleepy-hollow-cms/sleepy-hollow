package gateway_test

import (
	"content-management-api/domain"
	"content-management-api/driver/model"
	"content-management-api/gateway"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntry(t *testing.T) {

	target := gateway.Entry{}

	t.Run("Entryを登録しそのIDを返す", func(t *testing.T) {
		mockEntryDriver := new(MockContentDriver)

		inputEntry := domain.Entry{
			ContentModelID: "modelId",
		}

		returnEntry := model.Entry{
			ID:      "entryId",
			ModelID: "modelId",
		}
		inputModelEntry := model.Entry{
			ModelID: "modelId",
		}

		mockEntryDriver.On("CreateEntry", inputModelEntry).Return(&returnEntry, nil)

		target.Driver = mockEntryDriver

		actual, err := target.Create(context.TODO(), inputEntry)

		expected := domain.Entry{
			ID:             domain.EntryId("entryId"),
			ContentModelID: "modelId",
		}

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("EntryItemsを登録する", func(t *testing.T) {
		value, _ := domain.FactoryValue(domain.Text, "テキスト")
		inputEntryItems := []domain.EntryItem{
			{
				Type:      domain.Text,
				FieldName: "fieldName",
				Value:     value,
			},
		}

		entryItems := []model.EntryItem{
			{
				Type:  "text",
				Name:  "fieldName",
				Value: "テキスト",
			},
		}

		mockEntryDriver := new(MockContentDriver)
		mockEntryDriver.On("CreateEntryItems", model.EntryID("id"), entryItems).Return(entryItems, nil)
		target.Driver = mockEntryDriver

		actual, err := target.CreateItems(context.TODO(), domain.EntryId("id"), inputEntryItems)

		expected := []domain.EntryItem{
			{
				FieldName: "fieldName",
				Type:      domain.Text,
				Value: domain.TextValue{
					Value: "テキスト",
				},
			},
		}

		mockEntryDriver.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Entryを取得する", func(t *testing.T) {

		returnEntry := model.Entry{
			ID:      "id",
			ModelID: "modelId",
			Items: []model.EntryItem{
				{
					Type:  "text",
					Name:  "title",
					Value: "タイトル",
				},
			},
		}

		mockEntryDriver := new(MockContentDriver)
		mockEntryDriver.On("FindEntryByID", "id").Return(&returnEntry, nil)

		target.Driver = mockEntryDriver

		expected := domain.Entry{
			ID:             domain.EntryId("id"),
			ContentModelID: domain.ContentModelID("modelId"),
			Items: []domain.EntryItem{
				{
					FieldName: "title",
					Type:      domain.Text,
					Value:     domain.TextValue{Value: "タイトル"},
				},
			},
		}

		id := domain.EntryId("id")
		actual, err := target.FindById(context.TODO(), id)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}
