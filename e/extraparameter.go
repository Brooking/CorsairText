package e

import (
	"fmt"
)

// ExtraParameterError indicates that a command was called with too many parameters
type ExtraParameterError struct {
	Command  string
	Actual   int
	Expected int
}

// Error returns a textual representation of the ExtraParameterError
func (e *ExtraParameterError) Error() string {
	return "extra parameter, expected " + fmt.Sprint(e.Expected) + ", actual " + fmt.Sprint(e.Actual)
}

// IsShowToUser marks this error to show the error to the user
func (e *ExtraParameterError) IsShowToUser() {}

// GetCommand marks this error to show a specific help screen
func (e *ExtraParameterError) GetCommand() string {
	return e.Command
}

// NewExtraParameterError creates a ExtraParameterError
func NewExtraParameterError(command string, expected int, actual int) error {
	return &ExtraParameterError{
		Command:  command,
		Actual:   actual,
		Expected: expected,
	}
}

// IsExtraParameterError indicated whether the error is a ExtraParameterError
func IsExtraParameterError(err error) bool {
	_, ok := err.(*ExtraParameterError)
	return ok
}
