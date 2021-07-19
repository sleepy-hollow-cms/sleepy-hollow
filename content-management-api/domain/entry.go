package domain

type EntryId string

type Entry struct {
	ID EntryId
}

func (c EntryId) String() string {
	return string(c)
}
