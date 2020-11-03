package base

import "corsairtext/action"

// Base is an abstraction of a space base
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Base interface {
	// Actions returns a list of possible actions at this base
	Actions() action.List
}

// NewBase creates a new Base
func NewBase(baseType Type) Base {
	switch baseType {
	case TypeNone:
		return nil
	case TypeDirt:
		return &base{
			actions: dirtActions,
		}
	case TypeFull:
		return &base{
			actions: fullActions,
		}
	default:
		return nil
	}
}

// base is the implementation of a Base
type base struct {
	actions action.List
}

// Actions returns a list of possible actions at this base
func (b *base) Actions() action.List {
	return b.actions
}

// Type enum describes what sort of base this is
type Type string

const (
	// TypeNone indicates that there is no base here
	TypeNone Type = ""

	// TypeDirt indicates that there is just a spot to land here
	TypeDirt Type = "dirt"

	// TypeFull indicates that there is a full base here
	TypeFull Type = "full"
)

// String returns a textual representation of Type
func (t *Type) String() string {
	return string(*t)
}

// dirtActions stores the possible actions for a dirt base
var dirtActions = action.List{
	action.TypeDig,
}

// fullActions stores the possible actions for a full base
var fullActions = action.List{
	action.TypeBuy,
	action.TypeSell,
}
