package domain

type FieldID string

func (n FieldID) String() string {
	return string(n)
}

type Field struct {
	Name     Name
	Type     Type
	Required Required
}

type Required bool

type Fields []Field
