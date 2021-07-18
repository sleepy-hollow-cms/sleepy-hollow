package field

type Field struct {
	Type     Type
	Required Required
}

type Required bool

type Fields []Field
