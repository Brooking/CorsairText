package universe

import (
	"corsairtext/action"
	"corsairtext/support"
	"corsairtext/universe/spot"
)

// Universe is the main data layer interface
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Universe interface {
	//
	// temporary methods, replace with a call to 'look'
	//
	Description() string
	Path() string
	Actions() action.List

	// Act is a command to do something
	Act() string
}

// NewUniverse creates a new Universe
func NewUniverse(s support.Support) Universe {
	u := &universe{
		s: s,
	}
	u.root, u.current = u.generateMap()
	return u
}

// universe is the concrete implimentation of Universe
type universe struct {
	s       support.Support
	root    spot.Spot
	current spot.Spot
	index   map[string]spot.Spot
}

func (u *universe) Description() string {
	return u.current.Description()
}

func (u *universe) Path() string {
	return u.current.Path()
}

func (u *universe) Actions() action.List {
	return u.current.Actions()
}
