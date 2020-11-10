package textui

import (
	"corsairtext/action"
	"corsairtext/e"
	"corsairtext/match"

	"github.com/pkg/errors"
)

// MakeCommandMatcher creates a text matcher seeded with commands
func MakeCommandMatcher() match.Matcher {
	var commands []string
	for name := range actionDescriptionMap {
		commands = append(commands, name)
	}
	return match.NewMatcher(commands, false)
}

// actionDescription gives information about an action
type actionDescription struct {
	actionType      action.Type
	ShortUsage      string
	Usage           string
	Name            string
	Parameters      []parameterType
	RequestExemplar interface{}
	ParseParameters func(*textUI, []string) (interface{}, error)
}

// actionDescriptionMap is the complete list of action descriptions
var actionDescriptionMap = map[string]*actionDescription{
	"help": {
		actionType:      action.TypeHelp,
		ShortUsage:      "help - List commands",
		Usage:           "help <command> - List command(s)",
		Name:            "help",
		Parameters:      []parameterType{parameterTypeOptAny},
		RequestExemplar: helpRequest{},
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0:
				return &helpRequest{}, nil
			case 1:
				return &helpRequest{Command: arg[0]}, nil
			default:
				return nil, e.NewExtraParameterError(action.TypeHelp, 1, len(arg))
			}
		},
	},
	"quit": {
		actionType:      action.TypeQuit,
		Usage:           "quit - Leave the game",
		Name:            "quit",
		RequestExemplar: quitRequest{},
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0:
				return &quitRequest{}, nil
			default:
				return nil, e.NewExtraParameterError(action.TypeQuit, 0, len(arg))
			}
		},
	},
	"look": {
		actionType:      action.TypeLook,
		Usage:           "look - Look around",
		Name:            "look",
		RequestExemplar: lookRequest{},
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0:
				return &lookRequest{}, nil
			default:
				return nil, e.NewExtraParameterError(action.TypeLook, 0, len(arg))
			}
		},
	},
	"go": {
		actionType:      action.TypeGo,
		ShortUsage:      "go   - Travel",
		Usage:           "go <destination> - Travel to destination",
		Name:            "go",
		Parameters:      []parameterType{parameterTypeOptAny},
		RequestExemplar: goRequest{},
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0:
				return &goRequest{}, nil
			case 1:
				return &goRequest{Destination: arg[0]}, nil
			default:
				return nil, e.NewExtraParameterError(action.TypeGo, 1, len(arg))
			}
		},
	},
	"dig": {
		actionType:      action.TypeDig,
		Usage:           "dig  - Mine for ore",
		Name:            "dig",
		RequestExemplar: digRequest{},
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0:
				return &digRequest{}, nil
			default:
				return nil, e.NewExtraParameterError(action.TypeDig, 0, len(arg))
			}
		},
	},
	"buy": {
		actionType:      action.TypeBuy,
		ShortUsage:      "buy  - Purchase items",
		Usage:           "buy <amount> <item> - Purchase specified amount of items",
		Name:            "buy",
		Parameters:      []parameterType{parameterTypeNumber, parameterTypeAny},
		RequestExemplar: buyRequest{},
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0, 1:
				return nil, e.NewMissingParameterError(action.TypeSell, 2, len(arg))
			case 2:
				amount, ok := t.parseNumber(arg[0])
				if !ok {
					return nil, errors.Errorf("internal: malformed parameter #1 (%v) of type %v", arg[0], action.TypeBuy)
				}
				return &buyRequest{Amount: amount, Item: arg[1]}, nil
			default:
				return nil, e.NewExtraParameterError(action.TypeBuy, 2, len(arg))
			}
		},
	},
	"sell": {
		actionType:      action.TypeSell,
		ShortUsage:      "sell - Sell items",
		Usage:           "sell <amount> <item> - Sell specified amount of items",
		Name:            "sell",
		Parameters:      []parameterType{parameterTypeNumber, parameterTypeAny},
		RequestExemplar: sellRequest{},
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0, 1:
				return nil, e.NewMissingParameterError(action.TypeSell, 2, len(arg))
			case 2:
				amount, ok := t.parseNumber(arg[0])
				if !ok {
					return nil, errors.Errorf("internal: malformed parameter #1 (%v) of type %v", arg[0], action.TypeSell)
				}
				return &sellRequest{Amount: amount, Item: arg[1]}, nil
			default:
				return nil, e.NewExtraParameterError(action.TypeSell, 2, len(arg))
			}
		},
	},
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

// parameterRegex provides the proper regex for parameter types
var parameterRegex = map[parameterType]string{
	parameterTypeNone:      `\b`,
	parameterTypeNumber:    `\b\d+\b`,
	parameterTypeAny:       `\b.+\b`,
	parameterTypeOptNumber: `\b\d+\b`,
	parameterTypeOptAny:    `\b.+\b`,
}
