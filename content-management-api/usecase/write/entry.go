package write

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
)

type Entry struct {
	ContentModelID domain.ContentModelID
	Items          []EntryItem
}

type EntryItem struct {
	ContentType field.Type
	FieldName   field.Name
	Value       interface{} // FIXME
}
