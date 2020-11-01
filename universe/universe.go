package universe

import (
	"corsairtext/support"
	"corsairtext/universe/action"
)

// Universe is the main data layer interface
type Universe interface {
	Actions() []action.ActionDescription
	Act(action.Action)
}

// NewUniverse creates a new Universe
func NewUniverse(s support.Support) Universe {
	return &universe{
		s: s,
	}
}

// universe is the concrete implimentation of Universe
type universe struct {
	s support.Support
}

// Actions returns a slice of actions for the current spot
func (u *universe) Actions() []action.ActionDescription {
	return []action.ActionDescription{}
}

// Act applies the action to the current spot
func (u *universe) Act(action.Action) {
}
