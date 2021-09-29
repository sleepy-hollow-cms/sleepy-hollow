package domain

type PublishedStatus bool

type EntryPublication struct {
	EntryId         EntryId
	PublishedStatus PublishedStatus
}
