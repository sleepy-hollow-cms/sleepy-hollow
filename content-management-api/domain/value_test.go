package domain_test

import (
	"content-management-api/domain"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextValue(t *testing.T) {
	t.Run("stringを渡せばTextValueを生成できる", func(t *testing.T) {
		actual, err := domain.NewTextValue("test")
		expected := domain.TextValue{Value: "test"}

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("nilを渡すとTextValueの生成に失敗する", func(t *testing.T) {
		_, err := domain.NewTextValue(nil)

		assert.NotNil(t, err)
	})

	t.Run("数値を渡すとTextValueの生成に失敗する", func(t *testing.T) {
		_, err := domain.NewTextValue(1)
		fmt.Println(err)
		assert.NotNil(t, err)
	})
}

func TestMultipleTextValue(t *testing.T) {
	t.Run("[]stringを渡せばMultipleTextValueを生成できる", func(t *testing.T) {
		actual, err := domain.NewMultipleTextValue([]string{"1", "2"})
		expected := domain.MultipleTextValue{Value: []string{"1", "2"}}
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("文字列interface{}のリストを渡すとMultipleTextValueを生成できる", func(t *testing.T) {
		actual, err := domain.NewMultipleTextValue(interface{}([]interface{}{"1", "2"}))
		expected := domain.MultipleTextValue{Value: []string{"1", "2"}}
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("数値interface{}のリストを渡すとMultipleTextValueの生成に失敗する", func(t *testing.T) {
		_, err := domain.NewMultipleTextValue(interface{}([]interface{}{1, 2}))
		fmt.Println(err)
		assert.NotNil(t, err)
	})

	t.Run("intのリストを渡すとMultipleTextValueの生成に失敗する", func(t *testing.T) {
		_, err := domain.NewMultipleTextValue(interface{}([]int{1, 1}))
		fmt.Println(err)
		assert.NotNil(t, err)
	})

	t.Run("nilを渡すとMultipleTextValueの生成に失敗する", func(t *testing.T) {
		_, err := domain.NewMultipleTextValue(nil)
		assert.NotNil(t, err)
	})
}

func TestNumberValue(t *testing.T) {
	t.Run("float64", func(t *testing.T) {
		var float64Value = float64(100.0)
		_, err := domain.NewNumberValue(float64Value)
		assert.Nil(t, err)
	})
	t.Run("int", func(t *testing.T) {
		var intValue = int(100)
		_, err := domain.NewNumberValue(intValue)
		assert.Nil(t, err)
	})
	t.Run("int64", func(t *testing.T) {
		var intValue = int64(100)
		_, err := domain.NewNumberValue(intValue)
		assert.Nil(t, err)
	})
	t.Run("float32", func(t *testing.T) {
		var float32Value = float32(100)
		_, err := domain.NewNumberValue(float32Value)
		assert.Nil(t, err)
	})
}

func TestFactoryValue(t *testing.T) {
	t.Run("type textを渡すとTextValueを生成することができる", func(t *testing.T) {
		actual, _ := domain.SupportValue(domain.Text, "てきすとです")

		expected := domain.TextValue{
			Value: "this is text",
		}

		assert.IsType(t, expected, actual)
	})

	t.Run("type multiple-textを渡すとMultipleTextValueを生成することができる valueが[]interface{}の場合", func(t *testing.T) {
		actual, _ := domain.SupportValue(domain.MultipleText, interface{}([]interface{}{"1", "2"}))

		expected := domain.MultipleTextValue{
			Value: []string{"1", "2"},
		}

		assert.IsType(t, expected, actual)
	})

	t.Run("type multiple-textを渡すとMultipleTextValueを生成することができる valueが[]stringの場合", func(t *testing.T) {
		actual, _ := domain.SupportValue(domain.MultipleText, interface{}([]string{"1", "2"}))

		expected := domain.MultipleTextValue{
			Value: []string{"1", "2"},
		}

		assert.IsType(t, expected, actual)
	})
}
