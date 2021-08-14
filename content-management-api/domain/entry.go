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

func (e Entry) Validate(contentModel ContentModel) error {

	var requiredFieldNum int
	for _, item := range contentModel.Fields {
		if item.Required == true {
			requiredFieldNum++
		}
	}

	maxFieldNum := len(contentModel.Fields)
	if len(e.Items) < requiredFieldNum || maxFieldNum < len(e.Items) {
		return NewEntryValidationError("Number of fields not match")
	}

	var index int
	for i, field := range contentModel.Fields {
		if e.Items[index].Type != field.Type {
			if field.Required || i+1 == maxFieldNum {
				return NewEntryValidationError("Form of Entry field not match to Content Model")
			} else {
				continue
			}
		}

		index++
		if len(e.Items) <= index {
			break
		}
	}

	return nil
}

type EntryNotFound error
