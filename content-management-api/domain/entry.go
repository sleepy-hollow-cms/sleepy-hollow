package domain

type EntryId string

func (c EntryId) String() string {
	return string(c)
}

type Entry struct {
	ID             EntryId
	ContentModelID ContentModelID
}

type EntryItem struct {
	FieldName Name
	Type      Type
	Value     Value
}

type Validator interface {
	Do(p Type, value Value) error
}

func (e EntryItem) validate(validator Validator) {
	validator.Do(e.Type, e.Value)
}
