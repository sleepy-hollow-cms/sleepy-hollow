package driver

import "fmt"

type CannotFindByIdError struct {
	ID string
}

type (
	ContentModelCannotFindByIdError struct {
		CannotFindByIdError
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
