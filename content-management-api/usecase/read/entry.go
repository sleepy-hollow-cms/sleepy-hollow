package read

import (
	"content-management-api/domain"
)

type Entry struct {
	ID             domain.EntryId
	ContentModelID domain.ContentModelID
	EntryItems     EntryItem
}

type Item struct {
	FieldName domain.Name
	Type      domain.Type
	Value     domain.Value
}

type EntryItem struct {
	ID    domain.ID
	Items []Item
}
