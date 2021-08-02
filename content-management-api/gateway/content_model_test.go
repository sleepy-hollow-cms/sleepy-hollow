package gateway_test

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
	"content-management-api/driver"
	"content-management-api/driver/model"
	"content-management-api/gateway"
	"content-management-api/usecase"
	"content-management-api/usecase/write"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestContentModel(t *testing.T) {

	var target gateway.ContentModel

	t.Run("ContentModelを登録できる", func(t *testing.T) {
		createdAt := time.Now()
		mockContentModelDriver := new(MockContentDriver)
		contentModel := model.ContentModel{
			ID:        "id",
			Name:      "name",
			CreatedAt: createdAt,
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

		mockContentModelDriver.On("CreateModel", "name", createdAt, fields).Return(&contentModel, nil)
		target.Driver = mockContentModelDriver

		model := write.ContentModel{
			Name:      "name",
			CreatedAt: domain.CreatedAt(createdAt),
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
			ID:        domain.ContentModelID("id"),
			Name:      domain.Name("name"),
			CreatedAt: domain.CreatedAt(createdAt),
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

		createdAt := time.Now()

		mockContentDriver := new(MockContentDriver)

		id := domain.ContentModelID("id")
		model := model.ContentModel{
			ID:   "id",
			Name: "name",
			CreatedAt: createdAt,
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
		mockContentDriver.On("FindContentModelByID", "id").Return(&model, nil)

		target.Driver = mockContentDriver

		expected := domain.ContentModel{
			ID:   id,
			Name: domain.Name("name"),
			CreatedAt: domain.CreatedAt(createdAt),
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

	t.Run("Spaceに紐づくContentModelを取得することができる", func(t *testing.T) {

		createdAt := time.Now()
		mockContentDriver := new(MockContentDriver)
		id := domain.SpaceID("id")
		models := []model.ContentModel{
			{
				ID:        "id0",
				Name:      "name0",
				CreatedAt: createdAt,
				Fields: []model.Field{
					{
						Name:     "fname00",
						Type:     "text",
						Required: true,
					},
					{
						Name:     "fname01",
						Type:     "multiple-text",
						Required: false,
					},
				},
			},
			{
				ID:        "id1",
				Name:      "name1",
				CreatedAt: createdAt,
				Fields: []model.Field{
					{
						Name:     "fname10",
						Type:     "text",
						Required: true,
					},
					{
						Name:     "fname11",
						Type:     "multiple-text",
						Required: false,
					},
				},
			},
		}

		mockContentDriver.On("FindContentModelBySpaceID", "id").Return(models, nil)

		target.Driver = mockContentDriver

		expected := domain.ContentModels{
			SpaceID: id,
			Models: []domain.ContentModel{
				{
					ID:        "id0",
					Name:      domain.Name("name0"),
					CreatedAt: domain.CreatedAt(createdAt),
					Fields: field.Fields{
						{
							Name:     field.Name("fname00"),
							Type:     field.Text,
							Required: field.Required(true),
						},
						{
							Name:     field.Name("fname01"),
							Type:     field.MultipleText,
							Required: field.Required(false),
						},
					},
				},
				{
					ID:        "id1",
					Name:      domain.Name("name1"),
					CreatedAt: domain.CreatedAt(createdAt),
					Fields: field.Fields{
						{
							Name:     field.Name("fname10"),
							Type:     field.Text,
							Required: field.Required(true),
						},
						{
							Name:     field.Name("fname11"),
							Type:     field.MultipleText,
							Required: field.Required(false),
						},
					},
				},
			},
		}

		actual, err := target.FindBySpaceID(context.TODO(), id)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("ContentModelをID指定で取得し、存在なかった場合Errorを返す", func(t *testing.T) {

		mockContentModelDriver := new(MockContentDriver)
		id := domain.ContentModelID("id")

		contentModelCannotFind := driver.NewContentModelCannotFindById("id")

		mockContentModelDriver.On("FindContentModelByID", "id").Return(&model.ContentModel{}, contentModelCannotFind)

		target.Driver = mockContentModelDriver

		_, err := target.FindByID(context.TODO(), id)
		expected := usecase.NewContentModelNotFoundError("Contents Not Found By Id: id")

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &expected))
	})

	t.Run("ContentModelをID指定で取得時に不明なエラーが発生したらそのまま返す", func(t *testing.T) {

		mockContentModelDriver := new(MockContentDriver)
		id := domain.ContentModelID("id")

		someError := errors.New("some error")

		mockContentModelDriver.On("FindContentModelByID", "id").Return(&model.ContentModel{}, someError)

		target.Driver = mockContentModelDriver

		_, err := target.FindByID(context.TODO(), id)

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &someError))
	})

	t.Run("ContentModelをID指定で削除することができる", func(t *testing.T) {

		mockContentModelDriver := new(MockContentDriver)
		id := domain.ContentModelID("id")

		mockContentModelDriver.On("DeleteContentModelByID", "id").Return(nil)

		target.Driver = mockContentModelDriver

		err := target.DeleteByID(context.TODO(), id)

		mockContentModelDriver.AssertExpectations(t)
		assert.Nil(t, err)
	})

}
