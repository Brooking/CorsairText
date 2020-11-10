package e

// QuitError indicates that we received a quit command
type QuitError struct{}

// Error returns a textual representation of the QuitError
func (e *QuitError) Error() string {
	return "quit has been called"
}

// NewQuitError creates a QuitError
func NewQuitError() error {
	return &QuitError{}
}

// IsQuitError indicated whether the error is a QuitError
func IsQuitError(err error) bool {
	_, ok := err.(*QuitError)
	return ok
}
