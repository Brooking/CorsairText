package action

// Type enum indicates the action
type Type string

const (
	// TypeNone is a pseudo action for coding convenience
	TypeNone Type = ""

	//
	// Universal
	//

	// TypeQuit will end the game
	TypeQuit Type = "quit"

	// TypeHelp is a plea for which actions are available
	TypeHelp Type = "help"

	// TypeLook asks for current location information
	TypeLook Type = "look"

	// TypeGo is a move request
	TypeGo Type = "go"

	// TypeGoList is a move destination list request
	TypeGoList Type = "golist"

	//
	// Dirt
	//

	// TypeDig is an order to dig for ore
	TypeDig Type = "dig"

	//
	// Base
	//

	// TypeBuy is a buy order
	TypeBuy Type = "buy"

	// TypeSell is a sell order
	TypeSell Type = "sell"
)

// String returns a string representation of a Type value
func (t *Type) String() string {
	return string(*t)
}
