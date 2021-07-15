package model

type ContentModel struct {
	ID     string
	Name   string
	Fields []Field
}

type Field struct {
	Type string
}
