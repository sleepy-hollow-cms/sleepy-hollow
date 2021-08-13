package domain

import (
	"errors"
	"fmt"
)

type TextValue struct {
	Value string
}

func (t TextValue) FieldValue() interface{} {
	return t.Value
}

type MultipleTextValue struct {
	Value []string
}

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

func (m MultipleTextValue) FieldValue() interface{} {
	return m.Value
}

type NumberValue uint64

type Value interface {
	FieldValue() interface{}
}

type HasName interface {
	FieldName() string
}

func FactoryValue(typeName Type, value interface{}) (Value, error) {
	switch typeName {
	case Text:
		return NewTextValue(value)
	case MultipleText:
		return NewMultipleTextValue(value)
	default:
		return nil, fmt.Errorf("type not supported arg: %T:%v ", typeName, typeName)
	}
}
