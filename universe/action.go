package universe

import (
	"corsairtext/action"

	"github.com/pkg/errors"
)

// Action is the way things get done
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Action interface {
	Buy(int, string) error
	Go(string) error
	Help() ([]action.Type, error)
	Look() (string, string, string, error)
	Mine() error
	Quit() error
	Sell(int, string) error
}

// Buy acquires a certain amount of a commodity at a base
func (u *universe) Buy(amount int, commodity string) error {
	return errors.New("buy not yet implemented")
}

// Go moves current to a new spot
func (u *universe) Go(destination string) error {
	return errors.New("go not yet implemented")
}

// Help returns the list of actions available at the current spot
func (u *universe) Help() ([]action.Type, error) {
	actions := universalActions
	return actions.Append(u.current.Actions()), nil
}

// Look returns information about the current spot
func (u *universe) Look() ( /* Name */ string /* Description */, string /* Path */, string, error) {
	return u.current.Name(), u.current.Description(), u.current.Path(), nil
}

// Mine digs for ore
func (u *universe) Mine() error {
	return errors.New("mine not yet implemented")
}

// Quit shut everything down
func (u *universe) Quit() error {
	return nil
}

// Sell sells a certain amount of a commodity at a base
func (u *universe) Sell(amount int, commodity string) error {
	return errors.New("buy not yet implemented")
}

// actions returns the legal actions for the current spot
func (u *universe) actions() action.List {
	actions := universalActions
	return actions.Append(u.current.Actions())
}

// universalActions lists the universal actions
var universalActions = action.List{
	action.TypeHelp,
	action.TypeQuit,
}
