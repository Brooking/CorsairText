package e

// UnknownCommandError indicates that the entered command was not recognized
type UnknownCommandError struct {
	Command string
}

// Error returns a textual representation of the UnknownCommandError
func (e *UnknownCommandError) Error() string {
	return e.Command + " was not recognized as a valid command"
}

// IsShowToUser marks this error to show the error to the user
func (e *UnknownCommandError) IsShowToUser() {}

// IsShowAllHelp marks this error to show a full help screen
func (e *UnknownCommandError) IsShowAllHelp() {}

// NewUnknownCommandError creates a UnknownCommandError
func NewUnknownCommandError(command string) error {
	return &UnknownCommandError{
		Command: command,
	}
}

// IsUnknownCommandError indicated whether the error is a UnknownCommandError
func IsUnknownCommandError(err error) bool {
	_, ok := err.(*UnknownCommandError)
	return ok
}
