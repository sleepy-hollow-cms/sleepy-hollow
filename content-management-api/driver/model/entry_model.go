package model

type Entry struct {
	ID      string
	ModelID string
}

type EntryID string

func (e EntryID) String() string {
	return string(e)
}

type EntryItem struct {
	Type  string
	Name  string
	Value interface{}
}
