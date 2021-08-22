package domain

import "fmt"

type EntryId string

func (c EntryId) String() string {
	return string(c)
}

type Entry struct {
	ID             EntryId
	ContentModelID ContentModelID
	Items          []EntryItem
}

func (e Entry) CompareToModel(contentModel ContentModel) error {
	if len(e.Items) != len(contentModel.Fields) {
		return NewEntryValidationError("Number of fields not match")
	}

	for i, field := range contentModel.Fields {
		if field.Required && e.Items[i].Value == nil {
			return NewEntryValidationError("EntryItem is required")
		}
		if !field.Required && e.Items[i].Value == nil {
			continue
		}
		_, err := field.Type.Validate(e.Items[i].Value)
		if err != nil {
			return NewEntryValidationError(fmt.Sprintf("EntryItem does not match ContentModel\n %s", err.Error()))
		}
	}

	return nil
}

type EntryNotFound error
