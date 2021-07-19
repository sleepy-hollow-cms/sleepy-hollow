package usecase_test

import (
	"content-management-api/domain"
	"content-management-api/usecase"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestEntry(t *testing.T) {
	var target = usecase.Entry{}

	t.Run("Entryを登録することができる", func(t *testing.T) {
		mockEntryPort := new(MockEntryPort)

		entry := domain.Entry{
			ID: domain.EntryId("id"),
		}
		mockEntryPort.On("Create").Return(entry)

		target.EntryPort = mockEntryPort

		expected := domain.Entry{
			ID: domain.EntryId("id"),
		}
		actual, err := target.Create()

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}

type MockEntryPort struct {
	mock.Mock
}

func (_m *MockEntryPort) Create(ctx context.Context) (domain.Entry, error) {
	ret := _m.Called()
	return ret.Get(0).(domain.Entry), nil
}
