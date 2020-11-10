package e

// BadParameterError indicates that there was an unrecognized parameter
type BadParameterError struct {
	Command   string
	Parameter string
}

// Error returns a textual representation of the BadParameterError
func (e *BadParameterError) Error() string {
	return e.Parameter + " was not recognized as a valid parameter"
}

// IsShowToUser marks this error to show the error to the user
func (e *BadParameterError) IsShowToUser() {}

// GetCommand marks this error to show a specific help screen
func (e *BadParameterError) GetCommand() string {
	return e.Command
}

// NewBadParameterError creates a BadParameterError
func NewBadParameterError(command string, parameter string) error {
	return &BadParameterError{
		Command:   command,
		Parameter: parameter,
	}
}

// IsBadParameterError indicated whether the error is a BadParameterError
func IsBadParameterError(err error) bool {
	_, ok := err.(*BadParameterError)
	return ok
}
