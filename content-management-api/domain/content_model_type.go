package domain

type ContentModelType int

const (
	Entry ContentModelType = iota
	Media
)

func (c ContentModelType) String() string {
	switch c {
	case Entry:
		return "entry"
	case Media:
		return "media"
	default:
		return "Unknown"
	}
}
