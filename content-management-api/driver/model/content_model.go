package model

type ContentModel struct {
	ID     string
	Name   string
	Fields []Field
}

type Field struct {
	Type     string
	Required bool
}

type Entry struct {
	ID string
}
