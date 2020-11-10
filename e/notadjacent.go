package e

// NotAdjacentError indicates that a go was prohibited at a spot
type NotAdjacentError struct {
	Current     string
	Destination string
}

// Error returns a textual representation of the NotAdjacentError
func (e *NotAdjacentError) Error() string {
	return e.Destination + " not adjacent to " + e.Current
}

// IsShowToUser marks this error to show the error to the user
func (e *NotAdjacentError) IsShowToUser() {}

// IsShowAdjacent marks this error to show an adjacency help screen
func (e *NotAdjacentError) IsShowAdjacent() {}

// NewNotAdjacentError creates a NotAdjacentError
func NewNotAdjacentError(current string, destination string) error {
	return &NotAdjacentError{
		Current:     current,
		Destination: destination,
	}
}

// IsNotAdjacentError indicated whether the error is a NotAdjacentError
func IsNotAdjacentError(err error) bool {
	_, ok := err.(*NotAdjacentError)
	return ok
}
