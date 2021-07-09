package field

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
