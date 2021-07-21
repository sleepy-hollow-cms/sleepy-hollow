package gateway_test

import (
	"content-management-api/domain"
	"content-management-api/driver/model"
	"content-management-api/gateway"
	"content-management-api/usecase/write"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestEntry(t *testing.T) {

	target := gateway.Entry{}

	t.Run("Entryを登録しそのIDを返す", func(t *testing.T) {
		mockEntryDriver := new(MockEntryDriver)

		inputEntry := write.Entry{
			ContentModelID: "modelId",
		}

		returnEntry := model.Entry{
			ID:      "id",
			ModelID: "modelId",
		}
		inputModelEntry := model.Entry{
			ModelID: "modelId",
		}

		mockEntryDriver.On("Create", inputModelEntry).Return(&returnEntry)

		target.Driver = mockEntryDriver

		actual, err := target.Create(context.TODO(), inputEntry)

		expected := domain.Entry{
			ID:             domain.EntryId("id"),
			ContentModelID: "modelId",
		}

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}

type MockEntryDriver struct {
	mock.Mock
}

func (_m *MockEntryDriver) Create(entry model.Entry) (*model.Entry, error) {
	ret := _m.Called(entry)
	return ret.Get(0).(*model.Entry), nil
}
