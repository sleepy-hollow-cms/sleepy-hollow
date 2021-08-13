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
	panic("implemented me")
}

type EntryNotFound error
