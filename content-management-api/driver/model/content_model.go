package model

import "time"

type ContentModel struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Fields    []Field
}

type Field struct {
	Name     string
	Type     string
	Required bool
}

type Entry struct {
	ID      string
	ModelID string
}
