package model

type ContentModel struct {
	ID     string
	Fields []Field
}

type Field struct {
	Type string
}
