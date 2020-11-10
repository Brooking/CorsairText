package textui

import (
	"corsairtext/e"
	"corsairtext/match"

	"github.com/pkg/errors"
)

// MakeCommandMatcher creates a text matcher seeded with commands
func MakeCommandMatcher() match.Matcher {
	var commands []string
	for name := range commandDescriptionMap {
		commands = append(commands, name)
	}
	return match.NewMatcher(commands, false)
}

// Command is an enum for commands
type Command string

var (
	CommandBuy  = "buy"
	CommandDig  = "dig"
	CommandGo   = "go"
	CommandHelp = "help"
	CommandLook = "look"
	CommandQuit = "quit"
	CommandSell = "sell"
)

// commandHelpOrder dictates the order in which the commands are shown in help
var commandHelpOrder = []string{
	CommandHelp,
	CommandQuit,
	CommandLook,
	CommandGo,
	CommandDig,
	CommandBuy,
	CommandSell,
}

// commandDescription gives information about a command
type commandDescription struct {
	ShortUsage      string
	Usage           string
	ParseParameters func(*textUI, []string) (interface{}, error)
}

// commandDescriptionMap is the complete list of command descriptions
var commandDescriptionMap = map[string]*commandDescription{
	CommandBuy: {
		ShortUsage: "buy  - Purchase items",
		Usage:      "buy <amount> <item> - Purchase specified amount of items",
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0, 1:
				return nil, e.NewMissingParameterError(CommandBuy, 2, len(arg))
			case 2:
				amount, ok := t.parseNumber(arg[0])
				if !ok {
					return nil, errors.Errorf("internal: malformed parameter #1 (%v) of type %v", arg[0], CommandBuy)
				}
				return &buyCommand{Amount: amount, Item: arg[1]}, nil
			default:
				return nil, e.NewExtraParameterError(CommandBuy, 2, len(arg))
			}
		},
	},
	CommandDig: {
		Usage: "dig  - Mine for ore",
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0:
				return &digCommand{}, nil
			default:
				return nil, e.NewExtraParameterError(CommandDig, 0, len(arg))
			}
		},
	},
	CommandGo: {
		ShortUsage: "go   - Travel",
		Usage:      "go <destination> - Travel to destination",
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0:
				return &goCommand{}, nil
			case 1:
				return &goCommand{Destination: arg[0]}, nil
			default:
				return nil, e.NewExtraParameterError(CommandGo, 1, len(arg))
			}
		},
	},
	CommandHelp: {
		ShortUsage: "help - List commands",
		Usage:      "help <command> - List command(s)",
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0:
				return &helpCommand{}, nil
			case 1:
				return &helpCommand{Command: arg[0]}, nil
			default:
				return nil, e.NewExtraParameterError(CommandHelp, 1, len(arg))
			}
		},
	},
	CommandLook: {
		Usage: "look - Look around",
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0:
				return &lookCommand{}, nil
			default:
				return nil, e.NewExtraParameterError(CommandLook, 0, len(arg))
			}
		},
	},
	CommandQuit: {
		Usage: "quit - Leave the game",
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0:
				return &quitCommand{}, nil
			default:
				return nil, e.NewExtraParameterError(CommandQuit, 0, len(arg))
			}
		},
	},
	CommandSell: {
		ShortUsage: "sell - Sell items",
		Usage:      "sell <amount> <item> - Sell specified amount of items",
		ParseParameters: func(t *textUI, arg []string) (interface{}, error) {
			switch len(arg) {
			case 0, 1:
				return nil, e.NewMissingParameterError(CommandSell, 2, len(arg))
			case 2:
				amount, ok := t.parseNumber(arg[0])
				if !ok {
					return nil, errors.Errorf("internal: malformed parameter #1 (%v) of type %v", arg[0], CommandSell)
				}
				return &sellCommand{Amount: amount, Item: arg[1]}, nil
			default:
				return nil, e.NewExtraParameterError(CommandSell, 2, len(arg))
			}
		},
	},
}
