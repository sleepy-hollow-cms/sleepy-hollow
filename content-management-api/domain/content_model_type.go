package domain

type ContentModelType int

const (
	EntryType ContentModelType = iota
	MediaType
)

func (c ContentModelType) String() string {
	switch c {
	case EntryType:
		return "entry"
	case MediaType:
		return "media"
	default:
		return "Unknown"
	}
}
