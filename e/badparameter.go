package e

import "corsairtext/action"

// BadParameterError indicates that there was an unrecognized parameter
type BadParameterError struct {
	ActionType action.Type
	Parameter  string
}

// Error returns a textual representation of the BadParameterError
func (e *BadParameterError) Error() string {
	return e.Parameter + " was not recognized as a valid parameter"
}

// IsShowToUser marks this error to show the error to the user
func (e *BadParameterError) IsShowToUser() {}

// GetActionType marks this error to show a specific help screen
func (e *BadParameterError) GetActionType() *action.Type {
	return &e.ActionType
}

// NewBadParameterError creates a BadParameterError
func NewBadParameterError(actionType action.Type, parameter string) error {
	return &BadParameterError{
		ActionType: actionType,
		Parameter:  parameter,
	}
}

// IsBadParameterError indicated whether the error is a BadParameterError
func IsBadParameterError(err error) bool {
	_, ok := err.(*BadParameterError)
	return ok
}
