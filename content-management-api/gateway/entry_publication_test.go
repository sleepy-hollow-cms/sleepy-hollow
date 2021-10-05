package gateway_test

import (
	"context"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/gateway"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntryPublication(t *testing.T) {

	target := gateway.EntryPublication{}

	t.Run("Entryを公開する", func(t *testing.T) {
		mockEntryDriver := new(MockContentDriver)

		entryId := "entryId"
		inputEntryPublication := domain.EntryPublication{
			EntryId:         domain.EntryId(entryId),
			PublishedStatus: true,
		}

		returnEntry := model.Entry{
			ID:      entryId,
			ModelID: "modelId",
			Publication: model.Publication{
				Status: true,
			},
		}
		inputModelEntry := model.Entry{
			ID: entryId,
			Publication: model.Publication{
				Status: true,
			},
		}

		mockEntryDriver.On("UpdateEntry", inputModelEntry).Return(&returnEntry, nil)

		target.Driver = mockEntryDriver

		err := target.Store(context.TODO(), inputEntryPublication)

		mockEntryDriver.AssertExpectations(t)
		assert.Nil(t, err)
	})
}
