package universe

import (
	"corsairtext/universe/base"
	"corsairtext/universe/spot"
	"strings"

	"github.com/pkg/errors"
)

// generateMap create a map of spots
// returns the root of the map and a starting spot
func (u *universe) generateMap() error {
	u.index = make(map[string]spot.Spot)

	// todo load from json
	all, err := u.addSpot(spot.Init{Support: u.s, Name: "Galaxy", Description: "the whole thing"})
	if err != nil {
		return err
	}
	sol, err := u.addSpot(spot.Init{Support: u.s, Name: "Sol", Description: "a system", Parent: all})
	if err != nil {
		return err
	}
	earth, err := u.addSpot(spot.Init{Support: u.s, Name: "Earth", Description: "a planet", Parent: sol})
	if err != nil {
		return err
	}
	wm, err := u.addSpot(spot.Init{Support: u.s, Name: "Winnemucca", Description: "a full base", BaseType: base.TypeFull, Parent: earth})
	if err != nil || wm == nil {
		return err
	}
	luna, err := u.addSpot(spot.Init{Support: u.s, Name: "Luna", Description: "a moon", Parent: sol})
	if err != nil {
		return err
	}
	tb, err := u.addSpot(spot.Init{Support: u.s, Name: "Tranquility", Description: "a rough base", BaseType: base.TypeDirt, Parent: luna})
	if err != nil || tb == nil {
		return err
	}

	u.root = all
	u.current = wm
	return nil
}

// addSpot creates a spot, adds it to the index, and links it to its parent
func (u *universe) addSpot(init spot.Init) (spot.Spot, error) {
	spot := spot.NewSpot(init)
	_, exists := u.index[strings.ToLower(init.Name)]
	if exists {
		return nil, errors.Errorf("internal: %v already exists", init.Name)
	}
	u.index[strings.ToLower(init.Name)] = spot
	if init.Parent != nil {
		init.Parent.AddChild(spot)
	}
	return spot, nil
}
