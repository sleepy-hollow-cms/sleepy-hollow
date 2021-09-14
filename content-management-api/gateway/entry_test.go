package gateway_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/gateway"
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
		inputEntryItems := []domain.EntryItem{
			{
				Value: "テキスト",
			},
		}

		entryItems := []model.EntryItem{
			{
				Value: "テキスト",
			},
		}

		mockEntryDriver := new(MockContentDriver)
		mockEntryDriver.On("CreateEntryItems", model.EntryID("id"), entryItems).Return(entryItems, nil)
		target.Driver = mockEntryDriver

		actual, err := target.CreateItems(context.TODO(), domain.EntryId("id"), inputEntryItems)

		expected := []domain.EntryItem{
			{
				Value: "テキスト",
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
					Value: "タイトル",
				},
			},
		}

		id := domain.EntryId("id")
		actual, err := target.FindById(context.TODO(), id)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("EntryをID指定で削除する", func(t *testing.T) {
		mockEntryDriver := new(MockContentDriver)
		mockEntryDriver.On("DeleteEntryByID", "id").Return(int64(1), nil)

		target.Driver = mockEntryDriver

		id := domain.EntryId("id")
		err := target.DeleteById(context.TODO(), id)

		assert.Nil(t, err)
	})
}
