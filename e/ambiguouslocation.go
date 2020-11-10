package e

// AmbiguousLocationError indicates that the entered location was not recognized
type AmbiguousLocationError struct {
	Location      string
	Possibilities []string
}

// Error returns a textual representation of the AmbiguousLocationError
func (e *AmbiguousLocationError) Error() string {
	result := e.Location + " was ambiguous. Could be:\n"
	for _, location := range e.Possibilities {
		result = result + location + "\n"
	}
	return result
}

// IsShowToUser marks this error to show the error to the user
func (e *AmbiguousLocationError) IsShowToUser() {}

// NewAmbiguousLocationError creates a UnknownCommandError
func NewAmbiguousLocationError(location string, possibilities []string) error {
	return &AmbiguousLocationError{
		Location:      location,
		Possibilities: possibilities,
	}
}

// IsAmbiguousLocationError indicated whether the error is a AmbiguousLocationError
func IsAmbiguousLocationError(err error) bool {
	_, ok := err.(*AmbiguousLocationError)
	return ok
}
