package universe

import (
	"corsairtext/support"
	"corsairtext/universe/spot"
)

// Universe is the main data layer interface
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Universe interface {
	// WhereAmI returns the current spot
	WhereAmI() spot.Spot
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
}

// WhereAmI returns the current location
func (u *universe) WhereAmI() spot.Spot {
	return u.current
}

// generateMap create a map of spots
// returns the root of the map and a starting spot
func (u *universe) generateMap() (spot.Spot, spot.Spot) {
	// todo load from json
	all := spot.NewSpot(spot.Init{Name: "Galaxy", Description: "The whole thing"})
	sol := spot.NewSpot(spot.Init{Name: "Sol", Description: "A system", Parent: all})
	earth := spot.NewSpot(spot.Init{Name: "Earth", Description: "A planet", Parent: sol})
	wm := spot.NewSpot(spot.Init{Name: "Winnemucca Base", Description: "A landside base", BaseType: spot.BaseTypeFull, Parent: earth})
	return all, wm
}
