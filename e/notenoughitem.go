package e

import "fmt"

// NotEnoughItemError indicates that an action was prohibited at a spot
type NotEnoughItemError struct {
	Item string
	Have int
	Want int
}

// Error returns a textual representation of the NotEnoughItemError
func (e *NotEnoughItemError) Error() string {
	return fmt.Sprintf("not enough %s, %d requested, but only have %d", e.Item, e.Want, e.Have)
}

// IsShowToUser marks this error to show the error to the user
func (e *NotEnoughItemError) IsShowToUser() {}

// IsShowMarket marks this error to show a market help screen
func (e *NotEnoughItemError) IsShowMarket() {}

// NewNotEnoughItemError creates a NotEnoughItemError
func NewNotEnoughItemError(item string, have int, want int) error {
	return &NotEnoughItemError{
		Item: item,
		Have: have,
		Want: want,
	}
}

// IsNotEnoughItemError indicated whether the error is a NotEnoughItemError
func IsNotEnoughItemError(err error) bool {
	_, ok := err.(*NotEnoughItemError)
	return ok
}
