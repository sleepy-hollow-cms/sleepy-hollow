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
)

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

func (n DonotMatchByFilterError) Error() string {
	return fmt.Sprintf("Contents Not Match By Filter")
}

func NewContentModelCannotUpdateError() ContentModelCannotUpdateError {
	return ContentModelCannotUpdateError{
		DonotMatchByFilterError{},
	}
}
