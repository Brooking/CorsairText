package e

import (
	"corsairtext/action"
	"fmt"
)

// ExtraParameterError indicates that an action was called with too many parameters
type ExtraParameterError struct {
	ActionType action.Type
	Actual     int
	Expected   int
}

// Error returns a textual representation of the ExtraParameterError
func (e *ExtraParameterError) Error() string {
	return "extra parameter, expected " + fmt.Sprint(e.Expected) + ", actual " + fmt.Sprint(e.Actual)
}

// IsShowToUser marks this error to show the error to the user
func (e *ExtraParameterError) IsShowToUser() {}

// GetActionType marks this error to show a specific help screen
func (e *ExtraParameterError) GetActionType() *action.Type {
	return &e.ActionType
}

// NewExtraParameterError creates a ExtraParameterError
func NewExtraParameterError(actionType action.Type, expected int, actual int) error {
	return &ExtraParameterError{
		ActionType: actionType,
		Actual:     actual,
		Expected:   expected,
	}
}

// IsExtraParameterError indicated whether the error is a ExtraParameterError
func IsExtraParameterError(err error) bool {
	_, ok := err.(*ExtraParameterError)
	return ok
}
