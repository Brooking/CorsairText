package universe

import (
	"corsairtext/e"
	"corsairtext/universe/action"
	"strings"

	"github.com/pkg/errors"
)

// Action is the way things get done
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Action interface {
	Buy(int, string) error
	Go(string) error
	//	Help() (action.List, error)
	Dig() error
	Sell(int, string) error
}

// Buy acquires a certain amount of a commodity at a base
func (u *universe) Buy(amount int, commodity string) error {
	if !u.allowed(action.TypeBuy) {
		return e.NewBadSpotError(u.current, action.TypeBuy.String())
	}
	return errors.New("buy not yet implemented")
}

// Go moves current to a new spot
func (u *universe) Go(destination string) error {
	if !u.allowed(action.TypeGo) {
		return e.NewBadSpotError(u.current, action.TypeGo.String())
	}

	_, ok := u.index[strings.ToLower(destination)]
	if !ok {
		return e.NewUnknownDestinationError(destination)
	}

	adjacencies := u.current.ListAdjacent()
	for _, adjacent := range adjacencies {
		if strings.ToLower(destination) == strings.ToLower(adjacent.Name()) {
			u.current = adjacent
			return nil
		}
	}
	return e.NewNotAdjacentError(u.current.Name(), destination)
}

// Help returns the list of actions available at the current spot
func (u *universe) Help() (action.List, error) {
	if !u.allowed(action.TypeHelp) {
		return nil, e.NewBadSpotError(u.current, action.TypeHelp.String())
	}
	return u.current.Actions(), nil
}

// Dig mines for ore
func (u *universe) Dig() error {
	if !u.allowed(action.TypeDig) {
		return e.NewBadSpotError(u.current, action.TypeDig.String())
	}
	return errors.New("dig not yet implemented")
}

// Sell sells a certain amount of a commodity at a base
func (u *universe) Sell(amount int, commodity string) error {
	if !u.allowed(action.TypeSell) {
		return e.NewBadSpotError(u.current, action.TypeSell.String())
	}
	return errors.New("sell not yet implemented")
}

// allowed checks whether the action is allowed at the current spot
func (u *universe) allowed(actionType action.Type) bool {
	return u.current.Actions().Includes(actionType)
}
