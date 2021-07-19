package gateway_test

import (
	"content-management-api/domain"
	"content-management-api/driver/model"
	"content-management-api/gateway"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestEntry(t *testing.T) {

	target := gateway.Entry{}

	t.Run("Entryを登録しそのIDを返す", func(t *testing.T) {
		mockEntryDriver := new(MockEntryDriver)

		entry := model.Entry{
			ID: "id",
		}
		mockEntryDriver.On("Create").Return(&entry)

		target.Driver = mockEntryDriver

		actual, err := target.Create(context.TODO())

		expected := domain.Entry{
			ID: domain.EntryId("id"),
		}

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}

type MockEntryDriver struct {
	mock.Mock
}

func (_m *MockEntryDriver) Create() (*model.Entry, error) {
	ret := _m.Called()
	return ret.Get(0).(*model.Entry), nil
}
