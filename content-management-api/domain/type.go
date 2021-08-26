package domain

import "fmt"

type Type int

const (
	Text Type = iota
	MultipleText
	Date
	Number
	Bool
	Reference
	RichText
)

func (t Type) Validate(value interface{}) (Value, error) {
	switch t {
	case Text:
		return NewTextValue(value)
	case MultipleText:
		return NewMultipleTextValue(value)
	case Number:
		return NewNumberValue(value)
	case Date:
		return NewDateValue(value)
	case Bool:
		return NewBoolValue(value)
	case RichText:
		// TODO
		return nil, nil
	default:
		return nil, fmt.Errorf("type not supported arg: %T:%v", t.String(), t)
	}
}

func Of(value string) Type {
	switch value {
	case Text.String():
		return Text
	case MultipleText.String():
		return MultipleText
	case Date.String():
		return Date
	case Number.String():
		return Number
	case Bool.String():
		return Bool
	case Reference.String():
		return Reference
	case RichText.String():
		return RichText
	default:
		return -1
	}
}

func (t Type) String() string {
	switch t {
	case Text:
		return "text"
	case MultipleText:
		return "multiple-text"
	case Date:
		return "date"
	case Number:
		return "number"
	case Bool:
		return "bool"
	case Reference:
		return "reference"
	case RichText:
		return "rich-text"
	default:
		return "Unknown"
	}
}
