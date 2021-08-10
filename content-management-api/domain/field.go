package domain

type ID string

func (n ID) String() string {
	return string(n)
}

type Field struct {
	Name     Name
	Type     Type
	Required Required
}

type Required bool

type Fields []Field
