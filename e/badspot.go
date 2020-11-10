package e

import (
	"corsairtext/universe/spot"
)

// BadSpotError indicates that an action was prohibited at a spot
type BadSpotError struct {
	Spot   spot.Spot
	Action string
}

// Error returns a textual representation of the BadSpotError
func (e *BadSpotError) Error() string {
	return e.Action + " not allowed at " + e.Spot.Name()
}

// IsShowToUser marks this error to show the error to the user
func (e *BadSpotError) IsShowToUser() {}

// IsShowAllHelp marks this error to show a full help screen
func (e *BadSpotError) IsShowAllHelp() {}

// NewBadSpotError creates a BadSpotError
func NewBadSpotError(spot spot.Spot, action string) error {
	return &BadSpotError{
		Spot:   spot,
		Action: action,
	}
}

// IsBadSpotError indicated whether the error is a BadSpotError
func IsBadSpotError(err error) bool {
	_, ok := err.(*BadSpotError)
	return ok
}
