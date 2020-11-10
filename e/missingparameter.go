package e

import (
	"fmt"
)

// MissingParameterError indicates that a command was called with too few parameters
type MissingParameterError struct {
	Command  string
	Actual   int
	Expected int
}

// Error returns a textual representation of the MissingParameterError
func (e *MissingParameterError) Error() string {
	return "missing parameter, expected " + fmt.Sprint(e.Expected) + ", actual " + fmt.Sprint(e.Actual)
}

// IsShowToUser marks this error to show the error to the user
func (e *MissingParameterError) IsShowToUser() {}

// GetCommand marks this error to show a specific help screen
func (e *MissingParameterError) GetCommand() string {
	return e.Command
}

// NewMissingParameterError creates a MissingParameterError
func NewMissingParameterError(command string, expected int, actual int) error {
	return &MissingParameterError{
		Command:  command,
		Actual:   actual,
		Expected: expected,
	}
}

// IsMissingParameterError indicated whether the error is a MissingParameterError
func IsMissingParameterError(err error) bool {
	_, ok := err.(*MissingParameterError)
	return ok
}
