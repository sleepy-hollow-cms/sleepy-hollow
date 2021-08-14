package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntry(t *testing.T) {

	t.Run("Entryの数がContentModelの必須のFieldのより多い場合Errorを返す", func(t *testing.T) {

		target := Entry{
			Items: []EntryItem{
				{
					Type: Text,
				},
				{
					Type: Text,
				},
				{
					Type: Text,
				},
				{
					Type: Text,
				},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{
					Type:     Text,
					Required: true,
				},
				{
					Type:     Text,
					Required: false,
				},
				{
					Type:     Text,
					Required: true,
				},
			},
		}

		expectError := NewEntryValidationError("Number of fields not match")
		err := target.Validate(contentModel)

		assert.NotNil(t, err)
		assert.Equal(t, expectError, err)
	})

	t.Run("Entryの数がContentModelの必須のFieldの数より少ない場合はErrorを返す", func(t *testing.T) {

		target := Entry{
			Items: []EntryItem{
				{
					Type: Text,
				},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{
					Type:     Text,
					Required: Required(true),
				},
				{
					Type:     Text,
					Required: Required(false),
				},
				{
					Type:     Text,
					Required: Required(true),
				},
			},
		}

		expectError := NewEntryValidationError("Number of fields not match")
		err := target.Validate(contentModel)

		assert.NotNil(t, err)
		assert.Equal(t, expectError, err)
	})

	t.Run("Entryの形がContentModelに沿っている場合はnilを返す", func(t *testing.T) {

		target := Entry{
			Items: []EntryItem{
				{
					Type: Text,
				},
				{
					Type: MultipleText,
				},
				{
					Type: Number,
				},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{
					Type:     Text,
					Required: true,
				},
				{
					Type:     MultipleText,
					Required: true,
				},
				{
					Type:     Number,
					Required: true,
				},
			},
		}

		err := target.Validate(contentModel)

		assert.Nil(t, err)
	})

	t.Run("Entryの形がContentModelに沿っていない場合はErrorをかえす", func(t *testing.T) {

		target := Entry{
			Items: []EntryItem{
				{
					Type: Text,
				},
				{
					Type: MultipleText,
				},
				{
					Type: Number,
				},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{
					Type:     Text,
					Required: true,
				},
				{
					Type:     MultipleText,
					Required: true,
				},
				{
					Type:     MultipleText,
					Required: true,
				},
			},
		}

		expectError := NewEntryValidationError("Form of Entry field not match to Content Model")
		err := target.Validate(contentModel)

		assert.NotNil(t, err)
		assert.Equal(t, expectError, err)
	})

	t.Run("Entryの形が必須フィールドを考慮してContentModelに沿っている場合はnilをかえす", func(t *testing.T) {

		target := Entry{
			Items: []EntryItem{
				{
					Type: Text,
				},
				{
					Type: Number,
				},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{
					Type:     Text,
					Required: true,
				},
				{
					Type:     MultipleText,
					Required: false,
				},
				{
					Type:     Number,
					Required: true,
				},
				{
					Type:     Reference,
					Required: false,
				},
			},
		}

		err := target.Validate(contentModel)

		assert.Nil(t, err)
	})

	t.Run("Entryの形が必須フィールドを考慮してContentModelに沿っている場合はnilをかえす", func(t *testing.T) {

		target := Entry{
			Items: []EntryItem{
				{
					Type: MultipleText,
				},
				{
					Type: Reference,
				},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{
					Type:     Text,
					Required: false,
				},
				{
					Type:     MultipleText,
					Required: true,
				},
				{
					Type:     Number,
					Required: false,
				},
				{
					Type:     Reference,
					Required: true,
				},
			},
		}

		err := target.Validate(contentModel)

		assert.Nil(t, err)
	})

	t.Run("Entryの形が必須フィールドを考慮してContentModelに沿っている場合はnilをかえす", func(t *testing.T) {

		target := Entry{
			Items: []EntryItem{
				{
					Type: Text,
				},
				{
					Type: MultipleText,
				},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{
					Type:     Text,
					Required: true,
				},
				{
					Type:     MultipleText,
					Required: true,
				},
				{
					Type:     Number,
					Required: false,
				},
				{
					Type:     Reference,
					Required: false,
				},
			},
		}

		err := target.Validate(contentModel)

		assert.Nil(t, err)
	})

	t.Run("Entryの形が必須フィールドを考慮してContentModelに沿っている場合はnilをかえす", func(t *testing.T) {

		target := Entry{
			Items: []EntryItem{
				{
					Type: Number,
				},
				{
					Type: Reference,
				},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{
					Type:     Text,
					Required: false,
				},
				{
					Type:     MultipleText,
					Required: false,
				},
				{
					Type:     Number,
					Required: true,
				},
				{
					Type:     Reference,
					Required: true,
				},
			},
		}

		err := target.Validate(contentModel)

		assert.Nil(t, err)
	})

	t.Run("Entryの形が必須フィールドを考慮してContentModelに沿っていない場合はErrorを返す", func(t *testing.T) {

		target := Entry{
			Items: []EntryItem{
				{
					Type: Text,
				},
				{
					Type: Number,
				},
				{
					Type: Date,
				},
			},
		}

		contentModel := ContentModel{
			Fields: Fields{
				{
					Type:     Text,
					Required: true,
				},
				{
					Type:     MultipleText,
					Required: false,
				},
				{
					Type:     Number,
					Required: true,
				},
				{
					Type:     Reference,
					Required: false,
				},
			},
		}

		expectError := NewEntryValidationError("Form of Entry field not match to Content Model")
		err := target.Validate(contentModel)

		assert.NotNil(t, err)
		assert.Equal(t, expectError, err)
	})

}
