package usecase_test

import (
	"content-management-api/domain"
	"content-management-api/usecase"
	"content-management-api/usecase/write"
	"context"
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

		inputEntry := write.Entry{
			ContentModelID: domain.ContentModelID("modelId"),
		}

		model := domain.ContentModel{}

		mockEntryPort.On("Create", inputEntry).Return(entry)
		mockContentModelPort.On("FindByID", modelID).Return(model, nil)

		target.EntryPort = mockEntryPort
		target.ContentModelPort = mockContentModelPort

		expected := domain.Entry{
			ID:             domain.EntryId("id"),
			ContentModelID: domain.ContentModelID("modelId"),
		}
		actual, err := target.Create(inputEntry)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}

type MockEntryPort struct {
	mock.Mock
}

func (_m *MockEntryPort) Create(ctx context.Context, entry write.Entry) (domain.Entry, error) {
	ret := _m.Called(entry)
	return ret.Get(0).(domain.Entry), nil
}
