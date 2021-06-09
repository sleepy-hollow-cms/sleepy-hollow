package usecase_test

import (
	"content-management-api/domain"
	"content-management-api/support"
	"content-management-api/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContentModel(t *testing.T) {

	var target = usecase.ContentModelUseCase{}

	t.Run("ContentModelをIDを使って取得することができる", func(t *testing.T) {
		id := domain.ContentModelID("ID")

		actual, err := target.FindContentModelByID(id)

		if err != nil {
			t.Fatalf(err.Error())
		}

		expected := &domain.ContentModel{
			ID: id,
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("ContentModelを登録することができる", func(t *testing.T) {
		id := domain.ContentModelID("ID")
		model := domain.ContentModel{
			ID: id,
		}

		actual, err := target.CreateContentModel(model)

		if err != nil {
			t.Fatalf(err.Error())
		}

		expected := &model

		assert.Equal(t, expected, actual)
	})

	t.Run("Spaceに紐づくContentModelを全て取得することができる", func(t *testing.T) {
		t.Fatalf(support.TODO("Test code needs to be implemented.").Error())
	})
}
