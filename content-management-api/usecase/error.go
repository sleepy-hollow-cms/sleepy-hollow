package usecase

import "fmt"

type (
	SpaceNotFoundError struct {
		NotFoundError
	}
	ContentModelNotFoundError struct {
		NotFoundError
	}
	ContentModelCreateFailedError struct {
		Reason string
	}
)

type NotFoundError struct {
	Reason string
}

func (c ContentModelCreateFailedError) Error() string {
	return fmt.Sprintf("Register Failed. %s", c.Reason)
}

func (c ContentModelNotFoundError) Error() string {
	return fmt.Sprintf("Content Model Not Found. %s", c.Reason)
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

func NewContentModelCreateFailedError(reason string) ContentModelCreateFailedError {
	return ContentModelCreateFailedError{
		Reason: reason,
	}
}
