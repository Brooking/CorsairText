package universe

import (
	"corsairtext/action"
	"corsairtext/support"
	"corsairtext/universe/spot"
)

// Universe is the main data layer interface
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Universe interface {
	// Actions returns the legal actions for the current spot
	Actions() action.List

	// Act is a command to do something
	Act(action.Request) (interface{}, error)
}

// NewUniverse creates a new Universe
func NewUniverse(s support.Support) Universe {
	u := &universe{
		s: s,
	}
	u.root, u.current, _ = u.generateMap()
	return u
}

// universe is the concrete implimentation of Universe
type universe struct {
	s       support.Support
	root    spot.Spot
	current spot.Spot
	index   map[string]spot.Spot
}
