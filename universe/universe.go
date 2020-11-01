package universe

import "corsairtext/universe/action"

// Universe is the main data layer interface
type Universe interface {
	Actions() []action.ActionDescription
	Act(action.Action)
}

// NewUniverse creates a new Universe
func NewUniverse() Universe {
	return &universe{}
}

// universe is the concrete implimentation of Universe
type universe struct {
}

// Actions returns a slice of actions for the current spot
func (u *universe) Actions() []action.ActionDescription {
	return []action.ActionDescription{}
}

// Act applies the action to the current spot
func (u *universe) Act(action.Action) {
}
