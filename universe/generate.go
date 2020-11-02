package universe

import (
	"corsairtext/universe/base"
	"corsairtext/universe/spot"

	"github.com/pkg/errors"
)

// generateMap create a map of spots
// returns the root of the map and a starting spot
func (u *universe) generateMap() (spot.Spot, spot.Spot) {
	u.index = make(map[string]spot.Spot)

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
	if init.Parent != nil {
		init.Parent.AddChild(spot)
	}
	return spot, nil
}
