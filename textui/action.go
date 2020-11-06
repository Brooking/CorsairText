package textui

import (
	"corsairtext/action"
	"corsairtext/match"
)

// actionDescription gives human readable information about an action
type actionDescription struct {
	Type       action.Type
	ShortUsage string
	Usage      string
	Name       string
	Parameters []parameterType
}

// actionDescriptionTable is the complete list of action descriptions
var actionDescriptionTable = []actionDescription{
	{
		Type:       action.TypeHelp,
		ShortUsage: "Help - List commands",
		Usage:      "Help <command> - List command(s)",
		Name:       "help",
		Parameters: []parameterType{parameterTypeOptAny},
	},
	{
		Type:  action.TypeQuit,
		Usage: "Quit - Leave the game",
		Name:  "quit",
	},
	{
		Type:  action.TypeLook,
		Usage: "Look - Look around",
		Name:  "look",
	},
	{
		Type:       action.TypeGo,
		ShortUsage: "Go   - Travel",
		Usage:      "Go <destination> - Travel to destination",
		Name:       "go",
		Parameters: []parameterType{parameterTypeOptAny},
	},
	{
		Type:  action.TypeDig,
		Usage: "Dig  - Mine for ore",
		Name:  "dig",
	},
	{
		Type:       action.TypeBuy,
		ShortUsage: "Buy  - Purchase items",
		Usage:      "Buy <amount> <item> - Purchase specified amount of items",
		Name:       "buy",
		Parameters: []parameterType{parameterTypeNumber, parameterTypeAny},
	},
	{
		Type:       action.TypeSell,
		ShortUsage: "Sell - Sell items",
		Usage:      "Sell <amount> <item> - Sell specified amount of items",
		Name:       "sell",
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

// parameterType enum describes a parameter
type parameterType string

const (
	parameterTypeNone      parameterType = ""
	parameterTypeNumber    parameterType = "ParameterTypeNumber"
	parameterTypeAny       parameterType = "ParameterTypeAny"
	parameterTypeOptNumber parameterType = "parameterTypeOptNumber"
	parameterTypeOptAny    parameterType = "parameterTypeOptAny"
)

// String returns a textual represention of a parameterType
func (p parameterType) String() string {
	return string(p)
}

func MakeCommandMatcher() match.Matcher {
	var commands []string
	for _, description := range actionDescriptionTable {
		commands = append(commands, description.Name)
	}
	return match.NewMatcher(commands, false)
}
