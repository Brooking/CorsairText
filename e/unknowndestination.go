package e

// UnknownDestinationError indicates that an action was prohibited at a spot
type UnknownDestinationError struct {
	Destination string
}

// Error returns a textual representation of the UnknownDestinationError
func (e *UnknownDestinationError) Error() string {
	return e.Destination + " not found"
}

// IsShowToUser marks this error to show the error to the user
func (e *UnknownDestinationError) IsShowToUser() {}

// IsShowAdjacent marks this error to show an adjacency help screen
func (e *UnknownDestinationError) IsShowAdjacent() {}

// NewUnknownDestinationError creates a UnknownDestinationError
func NewUnknownDestinationError(destination string) error {
	return &UnknownDestinationError{
		Destination: destination,
	}
}

// IsUnknownDestinationError indicated whether the error is a UnknownDestinationError
func IsUnknownDestinationError(err error) bool {
	_, ok := err.(*UnknownDestinationError)
	return ok
}
