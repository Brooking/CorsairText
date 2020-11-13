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
	Dig() error
	Sell(int, string) error
}

// Buy acquires a certain amount of a commodity at a base
func (u *universe) Buy(amount int, commodity string) error {
	if !u.allowed(action.TypeBuy) {
		return e.NewBadSpotError(u.current, action.TypeBuy.String())
	}

	// todo - add a market (supply, buy price, sell price)
	return errors.New("buy not yet implemented")
}

// Go moves current to a new spot
func (u *universe) Go(destination string) error {
	if !u.allowed(action.TypeGo) {
		return e.NewBadSpotError(u.current, action.TypeGo.String())
	}

	_, ok := u.index[strings.ToLower(destination)]
	if !ok {
		return e.NewNoItemRoomError(destination)
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

// Dig mines for ore
func (u *universe) Dig() error {
	if !u.allowed(action.TypeDig) {
		return e.NewBadSpotError(u.current, action.TypeDig.String())
	}

	if u.ship.ItemCapacity <= u.ship.Load() {
		return e.NewNoItemRoomError("ore")
	}
	itemLot, ok := u.ship.Items["ore"]
	if !ok {
		u.ship.Items["ore"] = &ItemLot{}
		itemLot = u.ship.Items["ore"]
	}
	itemLot.Count++
	return nil
}

// Sell sells a certain amount of a commodity at a base
func (u *universe) Sell(amount int, commodity string) error {
	if !u.allowed(action.TypeSell) {
		return e.NewBadSpotError(u.current, action.TypeSell.String())
	}

	itemLot, ok := u.ship.Items[commodity]
	if !ok {
		return e.NewNotEnoughItemError(commodity, 0, amount)
	}
	if itemLot.Count < amount {
		return e.NewNotEnoughItemError(commodity, itemLot.Count, amount)
	}

	// todo - add a market (supply, buy price, sell price, unwanted items)
	u.ship.Money += amount // * unitCost
	itemLot.Count -= amount
	if itemLot.Count == 0 {
		delete(u.ship.Items, commodity)
	}
	return nil
}

// allowed checks whether the action is allowed at the current spot
func (u *universe) allowed(actionType action.Type) bool {
	return u.current.Actions().Includes(actionType)
}
