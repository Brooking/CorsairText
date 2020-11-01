package action

type Type int

const (
	// Universal
	TypeLook Type = iota
	TypeHelp Type = iota

	// Dirt
	TypeMine Type = iota

	// Base
	TypeBuy  Type = iota
	TypeSell Type = iota
)

type List []Type

func (l List) Append(more List) List {
	for _, item := range more {
		l = append(l, item)
	}
	return l
}

type Description struct {
	Type        Type
	Description string
	Usage       string
	Shortcut    string
}

var table = map[Type]Description{
	TypeLook: {
		Type:        TypeLook,
		Description: "Look around",
		Usage:       "(L)ook",
		Shortcut:    "look",
	},
	TypeHelp: {
		Type:        TypeHelp,
		Description: "List possible commands",
		Usage:       "(H)elp",
		Shortcut:    "help",
	},
	TypeMine: {
		Type:        TypeMine,
		Description: "Dig for ore",
		Usage:       "(M)ine",
		Shortcut:    "mine",
	},
	TypeBuy: {
		Type:        TypeBuy,
		Description: "Purchase a commodity",
		Usage:       "(B)uy",
		Shortcut:    "buy",
	},
	TypeSell: {
		Type:        TypeSell,
		Description: "Sell a commodity",
		Usage:       "(S)ell",
		Shortcut:    "sell",
	},
}

func ActionDescriptions(list List) []Description {
	var descriptions []Description
	for _, actionType := range list {
		description, ok := table[actionType]
		if !ok {
			continue
		}
		descriptions = append(descriptions, description)
	}
	return descriptions
}
