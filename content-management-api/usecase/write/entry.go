package write

import (
	"content-management-api/domain"
	"content-management-api/domain/field"
)

type Entry struct {
	ContentModelID domain.ContentModelID
}

type EntryItem struct {
	FieldName field.Name
	Value     interface{} // FIXME
}
