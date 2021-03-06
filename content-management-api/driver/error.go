package driver

import "fmt"

type CannotFindByIdError struct {
	ID string
}

type DonotMatchByFilterError struct {
}

type (
	ContentModelCannotFindByIdError struct {
		CannotFindByIdError
	}
	ContentModelCannotUpdateError struct {
		DonotMatchByFilterError
	}
	EntryCannotUpdateError struct {
		DonotMatchByFilterError
	}
	EntryNotFoundError struct {
		CannotFindByIdError
	}
	UserNotFoundError struct {
		CannotFindByIdError
	}
	UserCannotUpdateError struct {
		DonotMatchByFilterError
	}
)

func (n EntryNotFoundError) Error() string {
	return fmt.Sprintf("Entry Not Found By Id: %s", n.ID)
}

func (n UserNotFoundError) Error() string {
	return fmt.Sprintf("User Not Found By Id: %s", n.ID)
}

func (n CannotFindByIdError) Error() string {
	return fmt.Sprintf("Contents Not Found By Id: %s", n.ID)
}

func NewContentModelCannotFindById(id string) ContentModelCannotFindByIdError {
	return ContentModelCannotFindByIdError{
		CannotFindByIdError{
			ID: id,
		},
	}
}

func NewEntryNotFound(id string) EntryNotFoundError {
	return EntryNotFoundError{
		CannotFindByIdError{
			ID: id,
		},
	}
}

func NewUserNotFound(id string) UserNotFoundError {
	return UserNotFoundError{
		CannotFindByIdError{
			ID: id,
		},
	}
}

func (n DonotMatchByFilterError) Error() string {
	return fmt.Sprintf("Contents Not Match By Filter")
}

func NewContentModelCannotUpdateError() ContentModelCannotUpdateError {
	return ContentModelCannotUpdateError{
		DonotMatchByFilterError{},
	}
}

func NewEntryCannotUpdateError() EntryCannotUpdateError {
	return EntryCannotUpdateError{
		DonotMatchByFilterError{},
	}
}

func NewUserCannotUpdateError() UserCannotUpdateError {
	return UserCannotUpdateError{
		DonotMatchByFilterError{},
	}
}
