package write

import (
	"content-management-api/domain"
)

type Entry struct {
	ContentModelID domain.ContentModelID
}

type EntryItem struct {
	FieldName domain.Name
	Type      domain.Type
	Value     domain.Value
}

type Validator interface {
	Do(p domain.Type, value domain.Value) error
}

func (e EntryItem) validaTypete(validator Validator) {
	validator.Do(e.Type, e.Value)
}
