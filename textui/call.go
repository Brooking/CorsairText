package textui

import (
	"corsairtext/e"
	"strings"

	"github.com/pkg/errors"
)

func (t *textUI) call(command interface{}) (bool, error) {
	switch r := command.(type) {
	case *buyCommand:
		return t.buy(r)
	case *digCommand:
		return t.dig(r)
	case *goCommand:
		return t._go(r)
	case *helpCommand:
		return t.help(r)
	case *lookCommand:
		return t.look(r)
	case *quitCommand:
		return t.quit(r)
	case *sellCommand:
		return t.sell(r)
	default:
		return false, errors.Errorf("internal: no call handler for %T", r)
	}
}

// buy command describes a bus command
type buyCommand struct {
	Amount int
	Item   string
}

// buy handles a buy command
func (t *textUI) buy(command *buyCommand) (bool, error) {
	return false, t.a.Buy(command.Amount, command.Item)
}

// go command describes a go command
type goCommand struct {
	Destination string
}

// _go handles a go command
func (t *textUI) _go(command *goCommand) (bool, error) {
	switch {
	case command.Destination == "":
		adjacency := t.i.ListAdjacentLocations()
		for _, neighbor := range adjacency {
			t.s.Out.Println(neighbor)
		}
	default:
		err := t.a.Go(command.Destination)
		if err != nil {
			return false, err
		}

		t.look(&lookCommand{})
	}
	return false, nil
}

// help command describes a help command
type helpCommand struct {
	Command string
}

// help handles a help command
func (t *textUI) help(command *helpCommand) (bool, error) {
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
				return false, errors.Errorf("internal: unknown command %v", command)
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
			return false, e.NewUnknownCommandError(command.Command)
		case 1:
			description, ok := commandDescriptionMap[command.Command]
			if !ok {
				return false, e.NewUnknownCommandError(command.Command)
			}
			t.s.Out.Println(description.Usage)
			return false, nil
		default:
			return false, e.NewUnknownCommandError(command.Command)
		}
	}
	return false, nil
}

// lookCommand describes a look command
type lookCommand struct{}

// look handles a look command
func (t *textUI) look(command *lookCommand) (bool, error) {
	location := t.i.LocalLocation()
	t.s.Out.Println(strings.Join([]string{"You are at ", location.Name, ", ", location.Description, "."}, ""))

	var path string
	for _, spot := range location.Path {
		path = path + spot + "/"
	}
	t.s.Out.Println(path)
	return false, nil
}

// dig command describes a dis command
type digCommand struct{}

// dig handles a mining command
func (t *textUI) dig(command *digCommand) (bool, error) {
	return false, t.a.Dig()
}

// quitCommand describes a quit command
type quitCommand struct{}

// quit handles a quit command
func (t *textUI) quit(command *quitCommand) (bool, error) {
	return true, nil
}

// sellCommand describes a sell command
type sellCommand struct {
	Amount int
	Item   string
}

// sell handles a sell command
func (t *textUI) sell(command *sellCommand) (bool, error) {
	return false, t.a.Sell(command.Amount, command.Item)
}
