package gateway

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/driver/mongo"
	"content-management-api/usecase/write"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestContentModel(t *testing.T) {

	var target ContentModel

	t.Run("テキスト文字列のフィールドを含んだContentModelを登録できる", func(t *testing.T) {
		// Mock setting
		mockContentModelDriver := new(MockContentModelDriver)
		id := primitive.NewObjectID()
		contentModel := mongo.ContentModel{
			ID: id,
			Fields: []mongo.Field{
				{Type: "text"},
			},
		}

		mockContentModelDriver.On("Create", []string{"text"}).Return(&contentModel, nil)
		target.Driver = mockContentModelDriver

		model := write.ContentModel{
			Fields: field.Fields{
				{Type: field.Text},
			},
		}

		actual, err := target.Create(context.TODO(), model)

		expected := domain.ContentModel{
			ID: domain.ContentModelID(id.Hex()),
			Fields: field.Fields{
				{Type: field.Text},
			},
		}

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}

type MockContentModelDriver struct {
	mock.Mock
}

func (_m *MockContentModelDriver) Create(strings []string) (*mongo.ContentModel, error) {
	ret := _m.Called(strings)
	return ret.Get(0).(*mongo.ContentModel), ret.Error(1)
}
