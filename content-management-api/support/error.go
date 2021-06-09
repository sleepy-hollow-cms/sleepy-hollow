package support

import "fmt"

type NotImplementedError struct {
	reason string
}

func (n NotImplementedError) Error() string {
	return fmt.Errorf("an operation is not implemented: %s", n.reason).Error()
}

func NewNotImplementedError(reason string) NotImplementedError {
	return NotImplementedError{
		reason: reason,
	}
}

func TODO(reason string) error {
	return NewNotImplementedError(reason)
}
