package domain

import "fmt"

type (
	EntryValidationError struct {
		Reason string
	}
)

func (e EntryValidationError) Error() string {
	return fmt.Sprintf("Entry Validation Error. %s", e.Reason)
}

func NewEntryValidationError(reason string) EntryValidationError {
	return EntryValidationError{
		Reason: reason,
	}
}
