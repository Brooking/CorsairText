package textui

import (
	"corsairtext/e"
	"corsairtext/textui/match"
	"corsairtext/universe"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// call is a table for calling command handlers by type
func (t *textUI) call(command interface{}) error {
	switch r := command.(type) {
	case *buyCommand:
		return t.buy(r)
	case *digCommand:
		return t.dig(r)
	case *goCommand:
		return t._go(r)
	case *helpCommand:
		return t.help(r)
	case *inventoryCommand:
		return t.inventory(r)
	case *lookCommand:
		return t.look(r)
	case *mapCommand:
		return t._map(r)
	case *quitCommand:
		return t.quit(r)
	case *sellCommand:
		return t.sell(r)
	default:
		return errors.Errorf("internal: no call handler for %T", r)
	}
}

// buy command describes a bus command
type buyCommand struct {
	Amount int
	Item   string
}

// buy handles a buy command
func (t *textUI) buy(command *buyCommand) error {
	return t.a.Buy(command.Amount, command.Item)
}

// go command describes a go command
type goCommand struct {
	Destination string
}

// _go handles a go command
func (t *textUI) _go(command *goCommand) error {
	adjacency := t.i.ListAdjacentLocations()
	switch {
	case command.Destination == "":
		for _, neighbor := range adjacency {
			t.s.Out.Println(neighbor)
		}
	default:
		destinations := match.NewMatcher(adjacency, false).Match(command.Destination)
		switch len(destinations) {
		case 0:
			destinations = t.locationMatcher.Match(command.Destination)
			switch len(destinations) {
			case 0:
				return e.NewUnknownLocationError(command.Destination)
			case 1:
				return e.NewNotAdjacentError("", destinations[0])
			default:
				return e.NewAmbiguousLocationError(command.Destination, nil, destinations)
			}
		case 1:
			err := t.a.Go(destinations[0])
			if err != nil {
				return err
			}

			t.look(&lookCommand{})
		default:
			return e.NewAmbiguousLocationError(command.Destination, destinations, nil)
		}
	}
	return nil
}

// help command describes a help command
type helpCommand struct {
	Command string
}

// help handles a help command
func (t *textUI) help(command *helpCommand) error {
	switch {
	case command.Command == "":
		legalCommands := t.i.ListLocalCommands()
		for _, command := range commandHelpOrder {
			_, exist := legalCommands[command]
			if !exist {
				continue
			}
			description, ok := commandDescriptionMap[command]
			if !ok {
				return errors.Errorf("internal: unknown command %v", command)
			}
			usage := description.ShortUsage
			if usage == "" {
				usage = description.Usage
			}
			if usage == "" {
				continue
			}
			t.s.Out.Println(usage)
		}
	default:
		commands := t.commandMatcher.Match(command.Command)
		switch len(commands) {
		case 0:
			return e.NewUnknownCommandError(command.Command)
		case 1:
			description, ok := commandDescriptionMap[command.Command]
			if !ok {
				return e.NewUnknownCommandError(command.Command)
			}
			t.s.Out.Println(description.Usage)
			return nil
		default:
			return e.NewUnknownCommandError(command.Command)
		}
	}
	return nil
}

// inventoryCommand describes an inventory command
type inventoryCommand struct{}

// inventory handles an inventory command
func (t *textUI) inventory(command *inventoryCommand) error {
	inventory := t.i.Inventory()
	t.s.Out.Println(fmt.Sprintf("Money: %d", inventory.Money))
	t.s.Out.Println(fmt.Sprintf("Capacity: %d", inventory.ItemCapacity))
	t.s.Out.Println(fmt.Sprintf("Load: %d", inventory.Load()))
	for item, lot := range inventory.Items {
		t.s.Out.Println(fmt.Sprintf(" %s: %d", item, lot.Count))
	}
	return nil
}

// lookCommand describes a look command
type lookCommand struct{}

// look handles a look command
func (t *textUI) look(command *lookCommand) error {
	location := t.i.LocalLocation()
	t.s.Out.Println(strings.Join([]string{"You are at ", location.Name, ", ", location.Description, "."}, ""))

	var path string
	for _, spot := range location.Path {
		path = path + spot + "/"
	}
	t.s.Out.Println(path)
	return nil
}

// mapCommand describes a map command
type mapCommand struct {
	Location string
}

// _map handles a map command
func (t *textUI) _map(command *mapCommand) error {
	var name string
	locations := t.locationMatcher.Match(command.Location)
	if len(locations) == 1 {
		name = locations[0]
	}
	root := t.i.Map(name)
	t.mapDownWalker(root, 0)
	t.s.Out.Println("")
	t.mapUpWalker(root.Parent, root, 1)
	return nil
}

const indentBase = " "

// mapDownWalker is the recursive downward traverser for printing maps
func (t *textUI) mapDownWalker(node *universe.MapNode, depth int) {
	if node == nil {
		return
	}

	t.s.Out.Println(strings.Repeat(indentBase, depth) + node.Name)
	for _, child := range node.Children {
		t.mapDownWalker(child, depth+1)
	}
}

// mapUpWalker is the recursive upward traverser for printing maps
func (t *textUI) mapUpWalker(node *universe.MapNode, origin *universe.MapNode, depth int) {
	if node == nil {
		return
	}

	t.s.Out.Println(strings.Repeat(indentBase, depth) + node.Name)
	for _, child := range node.Children {
		if child == origin {
			continue
		}
		t.mapDownWalker(child, depth+1)
	}
	t.mapUpWalker(node.Parent, node, depth+1)
}

// dig command describes a dig command
type digCommand struct{}

// dig handles a mining command
func (t *textUI) dig(command *digCommand) error {
	return t.a.Dig()
}

// quitCommand describes a quit command
type quitCommand struct{}

// quit handles a quit command
func (t *textUI) quit(command *quitCommand) error {
	return e.NewQuitError()
}

// sellCommand describes a sell command
type sellCommand struct {
	Amount int
	Item   string
}

// sell handles a sell command
func (t *textUI) sell(command *sellCommand) error {
	return t.a.Sell(command.Amount, command.Item)
}
