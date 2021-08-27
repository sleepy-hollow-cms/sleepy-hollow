package domain

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntry(t *testing.T) {

	t.Run("Entryの数がContentModelのFieldと揃わない場合Errorを返す", func(t *testing.T) {
		target := Entry{
			Items: []EntryItem{
				{Value: "テキスト1"},
				{Value: "テキスト2"},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{Type: Text, Required: true},
				{Type: Text, Required: true},
				{Type: Text, Required: true},
			},
		}

		expectError := NewEntryValidationError("Number of fields not match")
		err := target.CompareToModel(contentModel)

		assert.NotNil(t, err)
		assert.Equal(t, expectError, err)
	})

	t.Run("Entryの数がContentModelのFieldと揃わない場合Errorを返す 全て必須じゃない場合でもErrorを返す", func(t *testing.T) {
		target := Entry{
			Items: []EntryItem{
				{Value: "テキスト1"},
				{Value: "テキスト2"},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{Type: Text, Required: false},
				{Type: Text, Required: false},
				{Type: Text, Required: false},
			},
		}

		expectError := NewEntryValidationError("Number of fields not match")
		err := target.CompareToModel(contentModel)

		assert.NotNil(t, err)
		assert.Equal(t, expectError, err)
	})
	t.Run("Entryの形がContentModelに沿っている場合はnilを返す", func(t *testing.T) {

		target := Entry{
			Items: []EntryItem{
				{Value: "テキスト"},
				{Value: "テキスト"},
				{Value: "テキスト"},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{Type: Text, Required: true},
				{Type: Text, Required: true},
				{Type: Text, Required: true},
			},
		}

		err := target.CompareToModel(contentModel)

		assert.Nil(t, err)
	})

	t.Run("Entryの形がContentModelに沿っていない場合はErrorを返す", func(t *testing.T) {

		target := Entry{
			Items: []EntryItem{
				{Value: "テキスト"},
				{Value: "テキスト"},
				{Value: "テキスト"},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{Type: Text, Required: true},
				{Type: Text, Required: true},
				{Type: MultipleText, Required: true},
			},
		}

		expectError := NewEntryValidationError("Form of Entry field not match to Content Model")
		err := target.CompareToModel(contentModel)

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &expectError))
	})

	t.Run("Entryの形がContentModelがサポートしていない型の場合はErrorを返す", func(t *testing.T) {

		target := Entry{
			Items: []EntryItem{
				{Value: []int{1, 2}},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{Type: Number, Required: true},
			},
		}

		expectError := NewEntryValidationError("error")
		err := target.CompareToModel(contentModel)

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &expectError))
	})

	t.Run("必須の型ではない場合Valueがnilでもerrorにならない", func(t *testing.T) {

		target := Entry{
			Items: []EntryItem{
				{Value: "テキスト"},
				{Value: nil},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{Type: Text, Required: true},
				{Type: Text, Required: false},
			},
		}

		err := target.CompareToModel(contentModel)

		assert.Nil(t, err)
	})

	t.Run("必須の型に対してValueがnilの場合error", func(t *testing.T) {
		target := Entry{
			Items: []EntryItem{
				{Value: "テキスト"},
				{Value: nil},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{Type: Text, Required: true},
				{Type: Text, Required: true},
			},
		}

		expectError := NewEntryValidationError("error")
		err := target.CompareToModel(contentModel)

		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &expectError))
	})

	t.Run("必須ではない型に対してValueがnilの場合errorにならない", func(t *testing.T) {
		target := Entry{
			Items: []EntryItem{
				{Value: nil},
				{Value: nil},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{Type: Text, Required: false},
				{Type: Text, Required: false},
			},
		}

		err := target.CompareToModel(contentModel)

		assert.Nil(t, err)
	})
}
