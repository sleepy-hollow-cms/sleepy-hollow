package usecase

import "fmt"

type (
	SpaceNotFoundError struct {
		NotFoundError
	}
	ContentModelNotFoundError struct {
		NotFoundError
	}
	EntryNotFoundError struct {
		NotFoundError
	}
	SpaceCreateFailedError struct {
		Reason string
	}
	SpaceUpdateFailedError struct {
		Reason string
	}
	ContentModelCreateFailedError struct {
		Reason string
	}
	ContentModelUpdateFailedError struct {
		Reason string
	}
	UserCreateFailedError struct {
		Reason string
	}
	ReferenceByEntryError struct {
		Reason string
	}
)

func (u ReferenceByEntryError) Error() string {
	return fmt.Sprintf("Delete Failed : %s", u.Reason)
}

func NewSpaceCreateFailedError(reason string) SpaceCreateFailedError {
	return SpaceCreateFailedError{Reason: reason}
}

func NewSpaceUpdateFailedError(reason string) SpaceUpdateFailedError {
	return SpaceUpdateFailedError{Reason: reason}
}

func NewUserCreateFailedError(reason string) UserCreateFailedError {
	return UserCreateFailedError{Reason: reason}
}

type NotFoundError struct {
	Reason string
}

func (c SpaceCreateFailedError) Error() string {
	return fmt.Sprintf("Register Failed. %s", c.Reason)
}

func (c SpaceUpdateFailedError) Error() string {
	return fmt.Sprintf("Update Failed. %s", c.Reason)
}

func (c ContentModelCreateFailedError) Error() string {
	return fmt.Sprintf("Register Failed. %s", c.Reason)
}

func (c ContentModelUpdateFailedError) Error() string {
	return fmt.Sprintf("Update Failed. %s", c.Reason)
}

func (c ContentModelNotFoundError) Error() string {
	return fmt.Sprintf("Content Model Not Found. %s", c.Reason)
}

func (c EntryNotFoundError) Error() string {
	return fmt.Sprintf("Entry Not Found. %s", c.Reason)
}

func (u UserCreateFailedError) Error() string {
	return fmt.Sprintf("Register Failed. %s", u.Reason)
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

func NewContentModelUpdateFailedError(reason string) ContentModelUpdateFailedError {
	return ContentModelUpdateFailedError{
		Reason: reason,
	}
}

func NewContentModelCreateFailedError(reason string) ContentModelCreateFailedError {
	return ContentModelCreateFailedError{
		Reason: reason,
	}
}

func NewEntryNotFoundError(reason string) EntryNotFoundError {
	return EntryNotFoundError{
		NotFoundError{
			Reason: reason,
		},
	}
}
