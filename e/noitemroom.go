package e

// NoItemRoomError indicates that an action was prohibited at a spot
type NoItemRoomError struct {
	Item string
}

// Error returns a textual representation of the NoItemRoomError
func (e *NoItemRoomError) Error() string {
	return "no room for " + e.Item
}

// IsShowToUser marks this error to show the error to the user
func (e *NoItemRoomError) IsShowToUser() {}

// IsShowMarket marks this error to show a market help screen
func (e *NoItemRoomError) IsShowMarket() {}

// NewNoItemRoomError creates a NoItemRoomError
func NewNoItemRoomError(location string) error {
	return &NoItemRoomError{
		Item: location,
	}
}

// IsNoItemRoomError indicated whether the error is a NoItemRoomError
func IsNoItemRoomError(err error) bool {
	_, ok := err.(*NoItemRoomError)
	return ok
}
