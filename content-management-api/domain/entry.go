package domain

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

		// 変換ができた場合はContentModelと合っている
		_, err := SupportValue(field.Type, e.Items[i].Value)
		if err != nil {
			return NewEntryValidationError("EntryItem does not match ContentModel")
		}
	}

	return nil
}

type EntryNotFound error
