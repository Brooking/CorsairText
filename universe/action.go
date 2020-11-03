package universe

import (
	"corsairtext/action"
	"corsairtext/e"

	"github.com/pkg/errors"
)

// Action is the way things get done
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Action interface {
	Buy(int, string) error
	Go(string) error
	Help() (action.List, error)
	Look() (*View, error)
	Dig() error
	Sell(int, string) error
}

// View is what you get when you look
type View struct {
	Name        string
	Description string
	Path        string
}

// Buy acquires a certain amount of a commodity at a base
func (u *universe) Buy(amount int, commodity string) error {
	if !u.allowed(action.TypeBuy) {
		return e.NewBadSpotError(u.current, action.TypeBuy)
	}
	return errors.New("buy not yet implemented")
}

// Go moves current to a new spot
func (u *universe) Go(destination string) error {
	if !u.allowed(action.TypeGo) {
		return e.NewBadSpotError(u.current, action.TypeGo)
	}
	return errors.New("go not yet implemented")
}

// Help returns the list of actions available at the current spot
func (u *universe) Help() (action.List, error) {
	if !u.allowed(action.TypeHelp) {
		return nil, e.NewBadSpotError(u.current, action.TypeHelp)
	}
	return u.current.Actions(), nil
}

// Look returns information about the current spot
func (u *universe) Look() (*View, error) {
	if !u.allowed(action.TypeLook) {
		return nil, e.NewBadSpotError(u.current, action.TypeLook)
	}
	return &View{
		Description: u.current.Description(),
		Name:        u.current.Name(),
		Path:        u.current.Path(),
	}, nil
}

// Dig mines for ore
func (u *universe) Dig() error {
	if !u.allowed(action.TypeDig) {
		return e.NewBadSpotError(u.current, action.TypeDig)
	}
	return errors.New("dig not yet implemented")
}

// Sell sells a certain amount of a commodity at a base
func (u *universe) Sell(amount int, commodity string) error {
	if !u.allowed(action.TypeSell) {
		return e.NewBadSpotError(u.current, action.TypeSell)
	}
	return errors.New("sell not yet implemented")
}

// allowed checks whether the action is allowed at the current spot
func (u *universe) allowed(actionType action.Type) bool {
	return u.current.Actions().Includes(actionType)
}
