package textui

import (
	"corsairtext/e"
	"corsairtext/textui/match"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// NewCommandMatcher creates a text matcher seeded with commands
func NewCommandMatcher() match.Matcher {
	var commands []string
	for name := range commandDescriptionMap {
		commands = append(commands, name)
	}
	return match.NewMatcher(commands, false)
}

// Command is an enum for commands
type Command string

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
	Handler    func(*textUI, []string) error
	ShortName  string
	LongName   string
	ShortUsage string
	LongUsage  string
}

// commandDescriptionMap is the complete list of command descriptions
var commandDescriptionMap = map[string]*commandDescription{
	CommandBuy: {
		Action:     true,
		ShortName:  "buy",
		LongName:   "buy <amount> <item>",
		ShortUsage: "Purchase items",
		LongUsage:  "Purchase specified amount of items",
		Handler: func(t *textUI, arg []string) error {
			switch len(arg) {
			case 0, 1:
				return e.NewMissingParameterError(CommandBuy, 2, len(arg))
			default:
				return e.NewExtraParameterError(CommandBuy, 2, len(arg))
			case 2:
				amount, ok := t.parseNumber(arg[0])
				if !ok {
					return errors.Errorf("internal: malformed parameter #1 (%v) of type %v", arg[0], CommandBuy)
				}
				return t.a.Buy(amount, arg[1])
			}
		},
	},
	CommandDig: {
		Action:     true,
		ShortName:  "dig",
		ShortUsage: "Mine for ore",
		Handler: func(t *textUI, arg []string) error {
			switch len(arg) {
			default:
				return e.NewExtraParameterError(CommandDig, 0, len(arg))
			case 0:
				return t.a.Dig()
			}
		},
	},
	CommandGo: {
		ShortName:  "go",
		LongName:   "go <destination>",
		ShortUsage: "Travel",
		LongUsage:  "Travel to destination",
		Handler: func(t *textUI, arg []string) error {
			switch len(arg) {
			default:
				return e.NewExtraParameterError(CommandGo, 1, len(arg))
			case 0:
				// show adjacent spots
				return t.showAdjacency()
			case 1:
				adjacency := t.i.ListAdjacentLocations()
				destinations := match.NewMatcher(adjacency, false).Match(arg[0])
				switch len(destinations) {
				default:
					return e.NewAmbiguousLocationError(arg[0], destinations, nil)
				case 0:
					// none found, try everywhere
					destinations = t.locationMatcher.Match(arg[0])
					switch len(destinations) {
					case 0:
						return e.NewUnknownLocationError(arg[0])
					case 1:
						return e.NewNotAdjacentError("", destinations[0])
					default:
						return e.NewAmbiguousLocationError(arg[0], nil, destinations)
					}
				case 1:
					err := t.a.Go(destinations[0])
					if err != nil {
						return err
					}

					t.showLook()
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
		Handler:    nil,
	},
	CommandInventory: {
		ShortName:  "inventory",
		ShortUsage: "Look at what you have",
		Handler: func(t *textUI, arg []string) error {
			switch len(arg) {
			default:
				return e.NewExtraParameterError(CommandInventory, 0, len(arg))
			case 0:
				inventory := t.i.Inventory()
				t.s.Out.Println(fmt.Sprintf("Money: %d", inventory.Money))
				t.s.Out.Println(fmt.Sprintf("Capacity: %d", inventory.ItemCapacity))
				t.s.Out.Println(fmt.Sprintf("Load: %d", inventory.Load()))
				for item, lot := range inventory.Items {
					t.s.Out.Println(fmt.Sprintf(" %s: %d", item, lot.Count))
				}
				return nil
			}
		},
	},
	CommandLook: {
		ShortName:  "look",
		ShortUsage: "Look around",
		Handler: func(t *textUI, arg []string) error {
			switch len(arg) {
			default:
				return e.NewExtraParameterError(CommandLook, 0, len(arg))
			case 0:
				return t.showLook()
			}
		},
	},
	CommandMap: {
		ShortName:  "map",
		LongName:   "map <location>",
		ShortUsage: "Show the map",
		LongUsage:  "Show the map around the location",
		Handler: func(t *textUI, arg []string) error {
			switch len(arg) {
			default:
				return e.NewExtraParameterError(CommandMap, 0, len(arg))
			case 0:
				return t.printMap(nil)
			case 1:
				var name *string
				locations := t.locationMatcher.Match(arg[0])
				if len(locations) == 1 {
					*name = locations[0]
				}
				return t.printMap(name)
			}
		},
	},
	CommandQuit: {
		ShortName:  "quit",
		ShortUsage: "Leave the game",
		Handler: func(t *textUI, arg []string) error {
			return e.NewQuitError()
		},
	},
	CommandSell: {
		Action:     true,
		ShortName:  "sell",
		LongName:   "sell <amount> <item>",
		ShortUsage: "Sell items",
		LongUsage:  "Sell specified amount of items",
		Handler: func(t *textUI, arg []string) error {
			switch len(arg) {
			case 0, 1:
				return e.NewMissingParameterError(CommandSell, 2, len(arg))
			default:
				return e.NewExtraParameterError(CommandSell, 2, len(arg))
			case 2:
				amount, ok := t.parseNumber(arg[0])
				if !ok {
					return errors.Errorf("internal: malformed parameter #1 (%v) of type %v", arg[0], CommandSell)
				}
				return t.a.Sell(amount, arg[1])
			}
		},
	},
}

// showLook implements the look command
func (t *textUI) showLook() error {
	location := t.i.LocalLocation()
	t.s.Out.Println(strings.Join([]string{"You are at ", location.Name, ", ", location.Description, "."}, ""))

	var path string
	for _, spot := range location.Path {
		path = path + spot + "/"
	}
	t.s.Out.Println(path)
	return nil
}

// showAllHelp implements the all help command
func (t *textUI) showAllHelp() error {
	var commands []*commandDescription
	var maxLen int
	legalActions := t.i.ListLocalActions()
	for _, command := range commandHelpOrder {
		description, ok := commandDescriptionMap[command]
		if !ok {
			return errors.Errorf("internal: unknown command %v", command)
		}
		if description.Action {
			_, exist := legalActions[command]
			if !exist {
				continue
			}
		}
		commands = append(commands, description)
		if len(description.ShortName) > maxLen {
			maxLen = len(description.ShortName)
		}
	}

	for _, description := range commands {
		name := description.ShortName + strings.Repeat(" ", maxLen-len(description.ShortName))
		t.s.Out.Println(strings.Join([]string{name, " - ", description.ShortUsage}, ""))
	}
	return nil
}

// showHelp implements the specific help command
func (t *textUI) showHelp(command string) error {
	commands := t.commandMatcher.Match(command)
	switch len(commands) {
	case 0:
		return e.NewUnknownCommandError(command)
	default:
		return e.NewUnknownCommandError(command)
	case 1:
		description, ok := commandDescriptionMap[command]
		if !ok {
			return e.NewUnknownCommandError(command)
		}

		name := description.LongName
		if name == "" {
			name = description.ShortName
		}
		usage := description.LongUsage
		if usage == "" {
			usage = description.ShortUsage
		}
		t.s.Out.Println(strings.Join([]string{name, " - ", usage}, ""))
		return nil
	}
}

// show adjacency shows the names of all adjacent spots
func (t *textUI) showAdjacency() error {
	adjacency := t.i.ListAdjacentLocations()
	for _, neighbor := range adjacency {
		t.s.Out.Println(neighbor)
	}
	return nil
}

// helpHandlerTableEntry is a replacement for the nil that is initially in the commandDescriptionMap
// go does not like initialization cycles, and so we initialize the table and this function seperately
// and then stick it in late (at textUI initialization)
func helpHandlerTableEntry(t *textUI, arg []string) error {
	switch len(arg) {
	default:
		return e.NewExtraParameterError(CommandHelp, 1, len(arg))
	case 0:
		return t.showAllHelp()
	case 1:
		return t.showHelp(arg[0])
	}
}
