package usecase

import "fmt"

type ContentModelNotFoundError NotFoundError
type ContentModelSaveFailError NotFoundError

func (c ContentModelSaveFailError) Error() string {
	return fmt.Sprintf("Save Failed. %s", c.Reason)
}

type SpaceNotFoundError NotFoundError

type NotFoundError struct {
	Reason string
}

func NewContentModelNotFoundError(reason string) ContentModelNotFoundError {
	return ContentModelNotFoundError{
		Reason: reason,
	}
}

func NewContentModelSaveFailError(reason string) ContentModelSaveFailError {
	return ContentModelSaveFailError{
		Reason: reason,
	}
}

func (n ContentModelNotFoundError) Error() string {
	return fmt.Sprintf("Not Found Error: %s", n.Reason)
}
