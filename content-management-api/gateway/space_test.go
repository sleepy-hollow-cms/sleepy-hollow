package gateway_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/driver/model"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/gateway"
)

func TestSpaceGateway(t *testing.T) {
	target := gateway.Space{}

	t.Run("Spaceを全て取得する", func(t *testing.T) {
		mockContentModelDriver := new(MockContentDriver)
		spaceModel := []model.Space{
			{ID: "spaceId", Name: "spaceName"},
		}
		mockContentModelDriver.On("FindSpace").Return(spaceModel, nil)
		target.Driver = mockContentModelDriver

		actual, err := target.Find(context.TODO())

		expected := domain.Spaces{
			{ID: domain.SpaceID("spaceId"), Name: domain.Name("spaceName")},
		}

		mockContentModelDriver.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("IDを指定してSpaceを取得する", func(t *testing.T) {
		id := domain.SpaceID("spaceId")

		mockContentModelDriver := new(MockContentDriver)
		spaceModel := model.Space{
			ID:   "spaceId",
			Name: "spaceName",
		}
		mockContentModelDriver.On("FindSpaceByID", id.String()).Return(&spaceModel, nil)

		target.Driver = mockContentModelDriver

		actual, err := target.FindByID(context.TODO(), id)

		expected := domain.Space{
			ID:   domain.SpaceID("spaceId"),
			Name: domain.Name("spaceName"),
		}

		mockContentModelDriver.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Spaceを作成する", func(t *testing.T) {
		mockContentModelDriver := new(MockContentDriver)
		createSpace := model.Space{
			Name: "spaceName",
		}
		mockContentModelDriver.On("CreateSpace", createSpace).Return(&model.Space{
			ID:   "spaceId",
			Name: "spaceName",
		}, nil)

		target.Driver = mockContentModelDriver

		actual, err := target.Register(context.TODO(), domain.Space{
			ID:   domain.SpaceID("spaceId"),
			Name: domain.Name("spaceName"),
		})

		expected := domain.Space{
			ID:   domain.SpaceID("spaceId"),
			Name: domain.Name("spaceName"),
		}

		mockContentModelDriver.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

}
