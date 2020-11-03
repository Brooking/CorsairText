package e

import (
	"corsairtext/action"
	"fmt"
)

// MissingParameterError indicates that an action was called with too few parameters
type MissingParameterError struct {
	ActionType action.Type
	Actual     int
	Expected   int
}

// Error returns a textual representation of the MissingParameterError
func (e *MissingParameterError) Error() string {
	return "missing parameter, expected " + fmt.Sprint(e.Expected) + ", actual " + fmt.Sprint(e.Actual)
}

// IsShowToUser marks this error to show the error to the user
func (e *MissingParameterError) IsShowToUser() {}

// GetActionType marks this error to show a specific help screen
func (e *MissingParameterError) GetActionType() *action.Type {
	return &e.ActionType
}

// NewMissingParameterError creates a MissingParameterError
func NewMissingParameterError(actionType action.Type, expected int, actual int) error {
	return &MissingParameterError{
		ActionType: actionType,
		Actual:     actual,
		Expected:   expected,
	}
}

// IsMissingParameterError indicated whether the error is a MissingParameterError
func IsMissingParameterError(err error) bool {
	_, ok := err.(*MissingParameterError)
	return ok
}
