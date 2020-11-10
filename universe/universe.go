package universe

import (
	"corsairtext/support"
	"corsairtext/universe/spot"
)

// NewUniverse creates a new Universe
func NewUniverse(s support.Support) (Action, Information) {
	u := &universe{
		s: s,
	}
	u.root, u.current, _ = u.generateMap()
	return u, u
}

// universe is the concrete implimentation of Universe
type universe struct {
	s       support.Support
	root    spot.Spot
	current spot.Spot
	index   map[string]spot.Spot
}
