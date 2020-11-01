package action

// Type indicates the action
type Type int

const (
	// Universal

	// TypeHelp is a plea for which actions are available
	TypeHelp Type = 0

	// TypeLook asks for current location information
	TypeLook Type = iota

	// TypeGo is a move request
	TypeGo Type = iota

	// Dirt

	// TypeMine is an order to dig for ore
	TypeMine Type = iota

	// Base

	// TypeBuy is a buy order
	TypeBuy Type = iota

	// TypeSell is a sell order
	TypeSell Type = iota
)

// List is a slice of action types
type List []Type

// Append allows two action lists to be concatinated
func (l List) Append(more List) List {
	for _, item := range more {
		l = append(l, item)
	}
	return l
}

// Description gives human readable information about an action
type Description struct {
	Type        Type
	Description string
	ShortUsage  string
	Usage       string
	Regex       string
}

// table is the complete list of action descriptions
var table = map[Type]Description{
	TypeHelp: {
		Type:        TypeHelp,
		Description: "List possible commands",
		ShortUsage:  "(H)elp",
		Usage:       "(H)elp",
		Regex:       "help",
	},
	TypeLook: {
		Type:        TypeLook,
		Description: "Look around",
		ShortUsage:  "(L)ook",
		Usage:       "(L)ook",
		Regex:       "look",
	},
	TypeGo: {
		Type:        TypeGo,
		Description: "Travel",
		ShortUsage:  "(G)o",
		Usage:       "(G)o (destination)",
		Regex:       "go",
	},
	TypeMine: {
		Type:        TypeMine,
		Description: "Dig for ore",
		ShortUsage:  "(M)ine",
		Usage:       "(M)ine",
		Regex:       "mine",
	},
	TypeBuy: {
		Type:        TypeBuy,
		Description: "Purchase a commodity",
		ShortUsage:  "(B)uy",
		Usage:       "(B)uy (number) (commodity)",
		Regex:       "buy",
	},
	TypeSell: {
		Type:        TypeSell,
		Description: "Sell a commodity",
		ShortUsage:  "(S)ell",
		Usage:       "(S)ell (number) (commodity)",
		Regex:       "sell",
	},
}

// Descriptions takes an action list and returns a slice of descriptions
func Descriptions(list List) []Description {
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
