package model

import "time"

type ContentModel struct {
	ID        string
	Name      string
	CreatedAt time.Time
	Fields    []Field
}

type Field struct {
	Name     string
	Type     string
	Required bool
}
