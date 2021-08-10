package domain

type Type int

const (
	Text Type = iota
	MultipleText
	Date
	Number
	Bool
	Reference
)

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
	default:
		return -1
	}
}

func (c Type) String() string {
	switch c {
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
	default:
		return "Unknown"
	}
}
