package domain

type EntryId string

type Entry struct {
	ID             EntryId
	ContentModelID ContentModelID
}

func (c EntryId) String() string {
	return string(c)
}
