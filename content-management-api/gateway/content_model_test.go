package gateway_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/gateway"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestContentModel(t *testing.T) {

	var target gateway.ContentModel

	t.Run("ContentModelを登録できる", func(t *testing.T) {
		createdAt := time.Now()

		mockContentModelDriver := new(MockContentDriver)

		contentModel := model.ContentModel{
			ID:        "contentmodelid",
			Name:      "name",
			CreatedAt: createdAt,
			UpdatedAt: createdAt,
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

		createdModel := domain.ContentModel{
			ID:        domain.ContentModelID("contentmodelid"),
			Name:      "name",
			CreatedAt: domain.CreatedAt(createdAt),
			UpdatedAt: domain.UpdatedAt(createdAt),
			Fields: domain.Fields{
				{
					Name:     domain.Name("fname"),
					Type:     domain.Text,
					Required: domain.Required(true),
				},
			},
		}

		actual, err := target.Create(context.TODO(), createdModel)

		expected := domain.ContentModel{
			ID:        domain.ContentModelID("contentmodelid"),
			Name:      domain.Name("name"),
			CreatedAt: domain.CreatedAt(createdAt),
			UpdatedAt: domain.UpdatedAt(createdAt),
			Fields: domain.Fields{
				{
					Name:     domain.Name("fname"),
					Type:     domain.Text,
					Required: domain.Required(true),
				},
			},
		}

		mockContentModelDriver.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("UpdatedAtが一致している時ContentModelを更新できる", func(t *testing.T) {

		createdAt := time.Now()
		updatedAt := time.Now()

		mockContentModelDriver := new(MockContentDriver)

		contentModel := model.ContentModel{
			ID:        "contentmodelid",
			Name:      "name",
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Fields: []model.Field{
				{
					Name:     "fname",
					Type:     "text",
					Required: true,
				},
			},
		}

		mockContentModelDriver.On("UpdateModel", mock.AnythingOfType("model.ContentModel")).Return(&contentModel, nil)
		target.Driver = mockContentModelDriver

		updatedModel := domain.ContentModel{
			ID:        domain.ContentModelID("contentmodelid"),
			Name:      "name",
			CreatedAt: domain.CreatedAt(createdAt),
			UpdatedAt: domain.UpdatedAt(updatedAt),
			Fields: domain.Fields{
				{
					Name:     domain.Name("fname"),
					Type:     domain.Text,
					Required: domain.Required(true),
				},
			},
		}

		actual, err := target.Update(context.TODO(), updatedModel)

		expected := domain.ContentModel{
			ID:        domain.ContentModelID("contentmodelid"),
			Name:      domain.Name("name"),
			CreatedAt: domain.CreatedAt(createdAt),
			UpdatedAt: domain.UpdatedAt(updatedAt),
			Fields: domain.Fields{
				{
					Name:     domain.Name("fname"),
					Type:     domain.Text,
					Required: domain.Required(true),
				},
			},
		}

		mockContentModelDriver.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("UpdatedAtが一致していない時ContentModelの更新ができずErrorが返る", func(t *testing.T) {

		createdAt := time.Now()
		updatedAt := time.Now()
		requestUpdatedAt := updatedAt.Add(time.Duration(1) * time.Second)

		mockContentModelDriver := new(MockContentDriver)

		mockContentModelDriver.On("UpdateModel", mock.AnythingOfType("model.ContentModel")).Return(&model.ContentModel{}, driver.NewContentModelCannotUpdateError())
		target.Driver = mockContentModelDriver

		updatedModel := domain.ContentModel{
			ID:        domain.ContentModelID("contentmodelid"),
			Name:      "name",
			CreatedAt: domain.CreatedAt(createdAt),
			UpdatedAt: domain.UpdatedAt(requestUpdatedAt),
			Fields: domain.Fields{
				{
					Name:     domain.Name("fname"),
					Type:     domain.Text,
					Required: domain.Required(true),
				},
			},
		}

		_, err := target.Update(context.TODO(), updatedModel)

		expected := usecase.NewContentModelUpdateFailedError("Content Model Update conflicted")

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &expected))
	})

	t.Run("ContentModelをID指定で取得することができる", func(t *testing.T) {

		createdAt := time.Now()
		updatedAt := time.Now()

		mockContentDriver := new(MockContentDriver)

		id := domain.ContentModelID("id")
		model := model.ContentModel{
			ID:        "id",
			Name:      "name",
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
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
			ID:        id,
			Name:      domain.Name("name"),
			CreatedAt: domain.CreatedAt(createdAt),
			UpdatedAt: domain.UpdatedAt(updatedAt),
			Fields: domain.Fields{
				{
					Name:     domain.Name("fname0"),
					Type:     domain.Text,
					Required: domain.Required(true),
				},
				{
					Name:     domain.Name("fname1"),
					Type:     domain.MultipleText,
					Required: domain.Required(false),
				},
			},
		}

		actual, err := target.FindByID(context.TODO(), id)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Spaceに紐づくContentModelを取得することができる", func(t *testing.T) {

		createdAt := time.Now()
		updatedAt := time.Now()
		mockContentDriver := new(MockContentDriver)
		id := domain.SpaceID("id")
		models := []model.ContentModel{
			{
				ID:        "id0",
				Name:      "name0",
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
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
				UpdatedAt: updatedAt,
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
					UpdatedAt: domain.UpdatedAt(updatedAt),
					Fields: domain.Fields{
						{
							Name:     domain.Name("fname00"),
							Type:     domain.Text,
							Required: domain.Required(true),
						},
						{
							Name:     domain.Name("fname01"),
							Type:     domain.MultipleText,
							Required: domain.Required(false),
						},
					},
				},
				{
					ID:        "id1",
					Name:      domain.Name("name1"),
					CreatedAt: domain.CreatedAt(createdAt),
					UpdatedAt: domain.UpdatedAt(updatedAt),
					Fields: domain.Fields{
						{
							Name:     domain.Name("fname10"),
							Type:     domain.Text,
							Required: domain.Required(true),
						},
						{
							Name:     domain.Name("fname11"),
							Type:     domain.MultipleText,
							Required: domain.Required(false),
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
