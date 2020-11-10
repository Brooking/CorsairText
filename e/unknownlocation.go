package e

// UnknownLocationError indicates that an action was prohibited at a spot
type UnknownLocationError struct {
	Location string
}

// Error returns a textual representation of the UnknownLocationError
func (e *UnknownLocationError) Error() string {
	return e.Location + " not found"
}

// IsShowToUser marks this error to show the error to the user
func (e *UnknownLocationError) IsShowToUser() {}

// IsShowAdjacent marks this error to show an adjacency help screen
func (e *UnknownLocationError) IsShowAdjacent() {}

// NewUnknownLocationError creates a UnknownDestinationError
func NewUnknownLocationError(location string) error {
	return &UnknownLocationError{
		Location: location,
	}
}

// IsUnknownLocationError indicated whether the error is a UnknownLocationError
func IsUnknownLocationError(err error) bool {
	_, ok := err.(*UnknownLocationError)
	return ok
}
