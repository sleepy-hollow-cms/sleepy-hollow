package gateway

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/driver/model"
	"content-management-api/usecase/write"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestContentModel(t *testing.T) {

	var target ContentModel

	t.Run("テキスト文字列のフィールドを含んだContentModelを登録できる", func(t *testing.T) {
		// Mock setting
		mockContentModelDriver := new(MockContentModelDriver)
		contentModel := model.ContentModel{
			ID:   "id",
			Name: "name",
			Fields: []model.Field{
				{
					Name:     "fname",
					Type:     "text",
					Required: true,
				},
			},
		}

		fields := []model.Field{
			{
				Name:     "fname",
				Type:     "text",
				Required: true,
			},
		}

		mockContentModelDriver.On("Create", "name", fields).Return(&contentModel, nil)
		target.Driver = mockContentModelDriver

		model := write.ContentModel{
			Name: "name",
			Fields: field.Fields{
				{
					Name:     field.Name("fname"),
					Type:     field.Text,
					Required: field.Required(true),
				},
			},
		}

		actual, err := target.Create(context.TODO(), model)

		expected := domain.ContentModel{
			ID:   domain.ContentModelID("id"),
			Name: domain.Name("name"),
			Fields: field.Fields{
				{
					Name:     field.Name("fname"),
					Type:     field.Text,
					Required: field.Required(true),
				},
			},
		}

		mockContentModelDriver.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("ContentModelをID指定で取得することができる", func(t *testing.T) {

		mockContentModelDriver := new(MockContentModelDriver)
		id := domain.ContentModelID("id")
		model := model.ContentModel{
			ID:   "id",
			Name: "name",
			Fields: []model.Field{
				{
					Name:     "fname0",
					Type:     "text",
					Required: true,
				},
				{
					Name:     "fname1",
					Type:     "multiple-text",
					Required: false,
				},
			},
		}
		mockContentModelDriver.On("FindByID", "id").Return(&model, nil)

		target.Driver = mockContentModelDriver

		expected := domain.ContentModel{
			ID:   id,
			Name: domain.Name("name"),
			Fields: field.Fields{
				{
					Name:     field.Name("fname0"),
					Type:     field.Text,
					Required: field.Required(true),
				},
				{
					Name:     field.Name("fname1"),
					Type:     field.MultipleText,
					Required: field.Required(false),
				},
			},
		}

		actual, err := target.FindByID(context.TODO(), id)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

}

type MockContentModelDriver struct {
	mock.Mock
}

func (_m *MockContentModelDriver) Create(name string, fields []model.Field) (*model.ContentModel, error) {
	ret := _m.Called(name, fields)
	return ret.Get(0).(*model.ContentModel), ret.Error(1)
}

func (_m *MockContentModelDriver) FindByID(id string) (*model.ContentModel, error) {
	ret := _m.Called(id)
	return ret.Get(0).(*model.ContentModel), ret.Error(1)
}
