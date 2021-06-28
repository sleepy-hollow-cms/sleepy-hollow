package usecase

import "fmt"

type (
	SpaceNotFoundError struct {
		NotFoundError
	}
	ContentModelNotFoundError struct {
		NotFoundError
	}
	ContentModelSaveFailError struct {
		Reason string
	}
)

type NotFoundError struct {
	Reason string
}

func (c ContentModelSaveFailError) Error() string {
	return fmt.Sprintf("Save Failed. %s", c.Reason)
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("Not Found Error: %s", n.Reason)
}

func NewSpaceNotFoundError(reason string) SpaceNotFoundError {
	return SpaceNotFoundError{
		NotFoundError{
			Reason: reason,
		},
	}
}

func NewContentModelNotFoundError(reason string) ContentModelNotFoundError {
	return ContentModelNotFoundError{
		NotFoundError{
			Reason: reason,
		},
	}
}

func NewContentModelSaveFailError(reason string) ContentModelSaveFailError {
	return ContentModelSaveFailError{
		Reason: reason,
	}
}
