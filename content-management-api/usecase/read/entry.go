package read

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
)

type Entry struct {
	ID             domain.EntryId
	ContentModelID domain.ContentModelID
	EntryItems     EntryItem
}

type Item struct {
	FieldName field.Name
	Type      field.Type
	Value     field.HasValue
}

type EntryItem struct {
	ID    field.ID
	Items []Item
}
