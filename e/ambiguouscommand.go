package e

// AmbiguousCommandError indicates that the entered command was not recognized
type AmbiguousCommandError struct {
	Command       string
	Possibilities []string
}

// Error returns a textual representation of the AmbiguousCommandError
func (e *AmbiguousCommandError) Error() string {
	result := e.Command + " was ambiguous. Could be:\n"
	for _, command := range e.Possibilities {
		result = result + command + "\n"
	}
	return result
}

// IsShowToUser marks this error to show the error to the user
func (e *AmbiguousCommandError) IsShowToUser() {}

// NewAmbiguousCommandError creates a UnknownCommandError
func NewAmbiguousCommandError(command string, possibilities []string) error {
	return &AmbiguousCommandError{
		Command:       command,
		Possibilities: possibilities,
	}
}

// IsAmbiguousCommandError indicated whether the error is a AmbiguousCommandError
func IsAmbiguousCommandError(err error) bool {
	_, ok := err.(*AmbiguousCommandError)
	return ok
}
