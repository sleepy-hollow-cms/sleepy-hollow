package domain

import (
	"errors"
	"fmt"
	"time"
)

type Value interface{}
type TextValue struct{ Value string }
type MultipleTextValue struct{ Value []string }
type NumberValue uint64
type DateValue struct{ Value time.Time }
type BoolValue struct{ Value bool }

func NewMultipleTextValue(value interface{}) (MultipleTextValue, error) {
	if value == nil {
		return MultipleTextValue{}, errors.New("failed factory MultipleTextValue. value is nil")
	}
	switch valueType := value.(type) {
	case []string:
		return MultipleTextValue{
			Value: value.([]string),
		}, nil
	case []interface{}:
		ir := value.([]interface{})
		strings := make([]string, len(ir))
		for i, v := range ir {
			switch vType := v.(type) {
			case string:
				strings[i] = v.(string)
			default:
				return MultipleTextValue{}, fmt.Errorf("Type mismatch error. interface conversion: interface {} is %T, not string. ", vType)
			}
		}
		return MultipleTextValue{
			Value: strings,
		}, nil
	default:
		return MultipleTextValue{}, fmt.Errorf("Type mismatch error. Only []string or []interface{} are allowed. arg: %T ", valueType)
	}
}

func NewTextValue(value interface{}) (TextValue, error) {
	if value == nil {
		return TextValue{}, errors.New("failed factory TextValue. value is nil")
	}
	switch valueType := value.(type) {
	case string:
		return TextValue{
			Value: value.(string),
		}, nil
	default:
		return TextValue{}, fmt.Errorf("Type mismatch error. Only string are allowed. arg: %T ", valueType)
	}
}

func NewNumberValue(value interface{}) (NumberValue, error) {
	switch valueType := value.(type) {
	case int:
		return NumberValue(value.(int)), nil
	case int64:
		return NumberValue(value.(int64)), nil
	case int32:
		return NumberValue(value.(int32)), nil
	case float32:
		return NumberValue(value.(float32)), nil
	case float64:
		return NumberValue(value.(float64)), nil
	default:
		return NumberValue(0), fmt.Errorf("Type mismatch error. Only int are allowed. arg: %T ", valueType)
	}
}

func NewBoolValue(value interface{}) (BoolValue, error) {
	switch valueType := value.(type) {
	case bool:
		return BoolValue{
			Value: value.(bool),
		}, nil
	default:
		return BoolValue{}, fmt.Errorf("Type mismatch error. Only bool are allowed. arg: %T ", valueType)
	}
}

func NewDateValue(value interface{}) (DateValue, error) {
	switch valueType := value.(type) {
	case string:
		parse, err := time.Parse(time.RFC3339, value.(string))
		if err != nil {
			return DateValue{}, err
		}
		return DateValue{Value: parse}, nil
	default:
		return DateValue{}, fmt.Errorf("Type mismatch error. Only date format are allowed. arg: %T ", valueType)
	}
}

func FactoryValue(value interface{}) (Value, error) {
	return value, nil
}
