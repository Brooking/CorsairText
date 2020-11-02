package universe

import (
	"corsairtext/support"
	"corsairtext/universe/base"
	"corsairtext/universe/spot"

	"github.com/pkg/errors"
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
	index   map[string]spot.Spot
}

// WhereAmI returns the current location
func (u *universe) WhereAmI() spot.Spot {
	return u.current
}

// generateMap create a map of spots
// returns the root of the map and a starting spot
func (u *universe) generateMap() (spot.Spot, spot.Spot) {
	// todo load from json
	all, _ := u.addSpot(spot.Init{Support: u.s, Name: "Galaxy", Description: "The whole thing"})
	sol, _ := u.addSpot(spot.Init{Support: u.s, Name: "Sol", Description: "A system", Parent: all})
	earth, _ := u.addSpot(spot.Init{Support: u.s, Name: "Earth", Description: "A planet", Parent: sol})
	wm, _ := u.addSpot(spot.Init{Support: u.s, Name: "Winnemucca Base", Description: "A landside base", BaseType: base.TypeFull, Parent: earth})
	return all, wm
}

// addSpot creates a spot, adds it to the index, and links it to its parent
func (u *universe) addSpot(init spot.Init) (spot.Spot, error) {
	spot := spot.NewSpot(init)
	_, exists := u.index[init.Name]
	if exists {
		return nil, errors.Errorf("%v already exists", init.Name)
	}
	u.index[init.Name] = spot
	init.Parent.AddChild(spot)
	return spot, nil
}
