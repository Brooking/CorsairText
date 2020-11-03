package textui

import "corsairtext/action"

// actionDescription gives human readable information about an action
type actionDescription struct {
	Type       action.Type
	ShortUsage string
	Usage      string
	NameRegex  string
	Parameters []parameterType
}

// actionDescriptionTable is the complete list of action descriptions
var actionDescriptionTable = []actionDescription{
	{
		Type:       action.TypeHelp,
		Usage:      "(H)elp - List possible commands",
		NameRegex:  "h(elp)?",
		Parameters: []parameterType{},
	},
	{
		Type:      action.TypeQuit,
		Usage:     "(Q)uit - Leave the game",
		NameRegex: "q(uit)?",
	},
	{
		Type:      action.TypeLook,
		Usage:     "(L)ook - Look around",
		NameRegex: "l(ook)?",
	},
	{
		Type:       action.TypeGo,
		ShortUsage: "(G)o - Travel",
		Usage:      "(G)o <destination> - Travel to destination",
		NameRegex:  "g(o)?",
		Parameters: []parameterType{parameterTypeAny},
	},
	{
		Type:      action.TypeMine,
		Usage:     "(M)ine - Dig for ore",
		NameRegex: "m(ine)?",
	},
	{
		Type:       action.TypeBuy,
		ShortUsage: "(B)uy - Purchase items",
		Usage:      "(B)uy <amount> <item> - Purchase specified amount of items",
		NameRegex:  "b(uy)?",
		Parameters: []parameterType{parameterTypeNumber, parameterTypeAny},
	},
	{
		Type:       action.TypeSell,
		ShortUsage: "(S)ell - Sell items",
		Usage:      "(S)ell <amount> <item> - Sell specified amount of items",
		NameRegex:  "s(ell)?",
		Parameters: []parameterType{parameterTypeNumber, parameterTypeAny},
	},
}

// describe returns a complete description of a Type
func describe(actionType action.Type) actionDescription {
	for _, description := range actionDescriptionTable {
		if description.Type != actionType {
			continue
		}
		return description
	}
	return actionDescription{}
}

// parameterType describes a parameter
type parameterType int

const (
	parameterTypeNone   parameterType = 0
	parameterTypeNumber parameterType = 1
	parameterTypeAny    parameterType = 2
)

// parameterToString is a mapping from ParameterType to a string describing it
var parameterTypeToString = map[parameterType]string{
	parameterTypeNone:   "ParameterTypeNone",
	parameterTypeNumber: "ParameterTypeNumber",
	parameterTypeAny:    "ParameterTypeAny",
}

// String returns a textual represention of a ParameterType
func (p parameterType) String() string {
	s, ok := parameterTypeToString[p]
	if !ok {
		return ""
	}
	return s
}
