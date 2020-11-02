package action

// Type indicates the action
type Type int

const (
	// TypeNone is a pseudo action for coding convenience
	TypeNone Type = 0

	//
	// Universal
	//

	// TypeQuit will end the game
	TypeQuit Type = iota

	// TypeHelp is a plea for which actions are available
	TypeHelp Type = iota

	// TypeLook asks for current location information
	TypeLook Type = iota

	// TypeGo is a move request
	TypeGo Type = iota

	//
	// Dirt
	//

	// TypeMine is an order to dig for ore
	TypeMine Type = iota

	//
	// Base
	//

	// TypeBuy is a buy order
	TypeBuy Type = iota

	// TypeSell is a sell order
	TypeSell Type = iota
)

// String returns a string representation of a Type value
func (t *Type) String() string {
	if t == nil {
		return ""
	}
	s, ok := typeToString[*t]
	if !ok {
		return ""
	}
	return s
}

// typeToString is a mapping between Types and strings
var typeToString = map[Type]string{
	TypeNone: "TypeNone",
	TypeQuit: "TypeQuit",
	TypeHelp: "TypeHelp",
	TypeLook: "TypeLook",
	TypeGo:   "TypeGo",
	TypeMine: "TypeMine",
	TypeBuy:  "TypeBuy",
	TypeSell: "TypeSell",
}
