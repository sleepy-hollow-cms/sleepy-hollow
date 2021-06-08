package domain

type ContentsModelID string

type ContentsModel struct {
	ID ContentsModelID
}

type ContentsModels struct {
	list []ContentsModel
}
