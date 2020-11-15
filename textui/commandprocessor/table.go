package commandprocessor

import (
	"corsairtext/e"
	"corsairtext/textui/commandprocessor/match"
	"fmt"

	"github.com/pkg/errors"
)

// NewCommandMatcher creates a text matcher seeded with commands
func NewCommandMatcher(cp *commandProcessor) match.Matcher {
	var commands []string
	for name := range cp.descriptions {
		commands = append(commands, name)
	}
	return match.NewMatcher(commands, false)
}

var (
	// CommandBuy the buy command
	CommandBuy = "buy"

	// CommandDig the dig command
	CommandDig = "dig"

	// CommandGo the go command
	CommandGo = "go"

	// CommandHelp the help command
	CommandHelp = "help"

	// CommandInventory the inventory command
	CommandInventory = "inventory"

	// CommandLook the look command
	CommandLook = "look"

	// CommandMap the map command
	CommandMap = "map"

	// CommandQuit the quit command
	CommandQuit = "quit"

	// CommandSell the sell command
	CommandSell = "sell"
)

// commandHelpOrder dictates the order in which the commands are shown in help
var commandHelpOrder = []string{
	CommandHelp,
	CommandQuit,
	CommandLook,
	CommandInventory,
	CommandMap,
	CommandGo,
	CommandDig,
	CommandBuy,
	CommandSell,
}

// commandDescription gives information about a command
type commandDescription struct {
	Action     bool
	Handler    func(*commandProcessor, []string) error
	ShortName  string
	LongName   string
	ShortUsage string
	LongUsage  string
}

// commandDescriptionMap is the complete list of command descriptions
func newDescriptions() map[string]*commandDescription {
	table := map[string]*commandDescription{
		CommandBuy: {
			Action:     true,
			ShortName:  "buy",
			LongName:   "buy <amount> <item>",
			ShortUsage: "Purchase items",
			LongUsage:  "Purchase specified amount of items",
			Handler: func(cp *commandProcessor, arg []string) error {
				switch len(arg) {
				case 0, 1:
					return e.NewMissingParameterError(CommandBuy, 2, len(arg))
				default:
					return e.NewExtraParameterError(CommandBuy, 2, len(arg))
				case 2:
					amount, ok := cp.parseNumber(arg[0])
					if !ok {
						return errors.Errorf("internal: malformed parameter #1 (%v) of type %v", arg[0], CommandBuy)
					}
					return cp.a.Buy(amount, arg[1])
				}
			},
		},
		CommandDig: {
			Action:     true,
			ShortName:  "dig",
			ShortUsage: "Mine for ore",
			Handler: func(cp *commandProcessor, arg []string) error {
				switch len(arg) {
				default:
					return e.NewExtraParameterError(CommandDig, 0, len(arg))
				case 0:
					return cp.a.Dig()
				}
			},
		},
		CommandGo: {
			ShortName:  "go",
			LongName:   "go <destination>",
			ShortUsage: "Travel",
			LongUsage:  "Travel to destination",
			Handler: func(cp *commandProcessor, arg []string) error {
				switch len(arg) {
				default:
					return e.NewExtraParameterError(CommandGo, 1, len(arg))
				case 0:
					// show adjacent spots
					return cp.ShowAdjacency()
				case 1:
					adjacency := cp.i.ListAdjacentLocations()
					destinations := match.NewMatcher(adjacency, false).Match(arg[0])
					switch len(destinations) {
					default:
						return e.NewAmbiguousLocationError(arg[0], destinations, nil)
					case 0:
						// none found, try everywhere
						destinations = cp.locationMatcher.Match(arg[0])
						switch len(destinations) {
						case 0:
							return e.NewUnknownLocationError(arg[0])
						case 1:
							return e.NewNotAdjacentError("", destinations[0])
						default:
							return e.NewAmbiguousLocationError(arg[0], nil, destinations)
						}
					case 1:
						err := cp.a.Go(destinations[0])
						if err != nil {
							return err
						}

						cp.ShowLook()
						return nil
					}
				}
			},
		},
		CommandHelp: {
			ShortName:  "help",
			LongName:   "help <command>",
			ShortUsage: "List commands",
			LongUsage:  "List command(s)",
			Handler: func(cp *commandProcessor, arg []string) error {
				switch len(arg) {
				default:
					return e.NewExtraParameterError(CommandHelp, 1, len(arg))
				case 0:
					return cp.ShowAllHelp()
				case 1:
					return cp.ShowHelp(arg[0])
				}
			},
		},
		CommandInventory: {
			ShortName:  "inventory",
			ShortUsage: "Look at what you have",
			Handler: func(cp *commandProcessor, arg []string) error {
				switch len(arg) {
				default:
					return e.NewExtraParameterError(CommandInventory, 0, len(arg))
				case 0:
					inventory := cp.i.Inventory()
					cp.s.Out.Println(fmt.Sprintf("Money: %d", inventory.Money))
					cp.s.Out.Println(fmt.Sprintf("Capacity: %d", inventory.ItemCapacity))
					cp.s.Out.Println(fmt.Sprintf("Load: %d", inventory.Load()))
					for item, lot := range inventory.Items {
						cp.s.Out.Println(fmt.Sprintf(" %s: %d", item, lot.Count))
					}
					return nil
				}
			},
		},
		CommandLook: {
			ShortName:  "look",
			ShortUsage: "Look around",
			Handler: func(cp *commandProcessor, arg []string) error {
				switch len(arg) {
				default:
					return e.NewExtraParameterError(CommandLook, 0, len(arg))
				case 0:
					return cp.ShowLook()
				}
			},
		},
		CommandMap: {
			ShortName:  "map",
			LongName:   "map <location>",
			ShortUsage: "Show the map",
			LongUsage:  "Show the map around the location",
			Handler: func(cp *commandProcessor, arg []string) error {
				switch len(arg) {
				default:
					return e.NewExtraParameterError(CommandMap, 0, len(arg))
				case 0:
					return cp.printMap(nil)
				case 1:
					var name *string
					locations := cp.locationMatcher.Match(arg[0])
					if len(locations) == 1 {
						*name = locations[0]
					}
					return cp.printMap(name)
				}
			},
		},
		CommandQuit: {
			ShortName:  "quit",
			ShortUsage: "Leave the game",
			Handler: func(cp *commandProcessor, arg []string) error {
				return e.NewQuitError()
			},
		},
		CommandSell: {
			Action:     true,
			ShortName:  "sell",
			LongName:   "sell <amount> <item>",
			ShortUsage: "Sell items",
			LongUsage:  "Sell specified amount of items",
			Handler: func(cp *commandProcessor, arg []string) error {
				switch len(arg) {
				case 0, 1:
					return e.NewMissingParameterError(CommandSell, 2, len(arg))
				default:
					return e.NewExtraParameterError(CommandSell, 2, len(arg))
				case 2:
					amount, ok := cp.parseNumber(arg[0])
					if !ok {
						return errors.Errorf("internal: malformed parameter #1 (%v) of type %v", arg[0], CommandSell)
					}
					return cp.a.Sell(amount, arg[1])
				}
			},
		},
	}
	// table[CommandHelp].Handler = helpHandlerTableEntry
	return table
}
