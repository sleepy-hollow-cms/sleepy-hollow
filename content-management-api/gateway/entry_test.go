package gateway_test

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/driver/model"
	"content-management-api/gateway"
	"content-management-api/usecase/read"
	"content-management-api/usecase/write"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntry(t *testing.T) {

	target := gateway.Entry{}

	t.Run("Entryを登録しそのIDを返す", func(t *testing.T) {
		mockEntryDriver := new(MockContentDriver)

		inputEntry := write.Entry{
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
		inputEntryItems := []write.EntryItem{
			{
				Type:      field.Text,
				FieldName: "fieldName",
				Value:     field.FactoryValue(field.Text, "テキスト"),
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

		items := []read.Item{
			{
				FieldName: "fieldName",
				Type:      field.Text,
				Value: field.TextValue{
					Value: "テキスト",
				},
			},
		}

		expected := read.EntryItem{
			ID:    field.ID("id"),
			Items: items,
		}

		mockEntryDriver.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}
