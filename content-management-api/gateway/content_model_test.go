package gateway_test

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/driver/model"
	"content-management-api/gateway"
	"content-management-api/usecase/write"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentModel(t *testing.T) {

	var target gateway.ContentModel

	t.Run("テキスト文字列のフィールドを含んだContentModelを登録できる", func(t *testing.T) {
		// Mock setting
		mockContentModelDriver := new(MockContentDriver)
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

		mockContentModelDriver.On("CreateModel", "name", fields).Return(&contentModel, nil)
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

}
