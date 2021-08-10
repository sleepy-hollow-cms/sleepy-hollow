package domain_test

import (
	field "content-management-api/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntry(t *testing.T) {
	t.Run("type textを渡すとTextValueを生成することができる", func(t *testing.T) {
		actual := field.FactoryValue(field.Text, "てきすとです")

		expected := field.TextValue{
			Value: "てきすとです",
		}

		assert.IsType(t, expected, actual)
	})

	t.Run("type multiple-textを渡すとMultipleTextValueを生成することができる valueが[]interface{}の場合", func(t *testing.T) {
		actual := field.FactoryValue(field.MultipleText, interface{}([]interface{}{"1", "2"}))

		expected := field.MultipleTextValue{
			Value: []string{"1", "2"},
		}

		assert.IsType(t, expected, actual)
	})

	t.Run("type multiple-textを渡すとMultipleTextValueを生成することができる valueが[]stringの場合", func(t *testing.T) {
		actual := field.FactoryValue(field.MultipleText, interface{}([]string{"1", "2"}))

		expected := field.MultipleTextValue{
			Value: []string{"1", "2"},
		}

		assert.IsType(t, expected, actual)
	})
}
