package model

type Entry struct {
	ID      string
	ModelID string
	Items   []EntryItem
}

type EntryID string

func (e EntryID) String() string {
	return string(e)
}

type EntryItem struct {
	Value interface{}
}
