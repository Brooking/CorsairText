package textui

import (
	"corsairtext/action"
	"strings"

	"github.com/pkg/errors"
)

func (t *textUI) call(request Request) (bool, error) {
	var actHandler = map[action.Type]func([]interface{}) (bool, error){
		action.TypeBuy:  t.buy,
		action.TypeGo:   t.move,
		action.TypeHelp: t.help,
		action.TypeLook: t.look,
		action.TypeMine: t.mine,
		action.TypeQuit: t.quit,
		action.TypeSell: t.sell,
	}
	handler, ok := actHandler[request.Type]
	if !ok {
		return false, errors.Errorf("no handler for %v", request.Type)
	}

	quit, err := handler(request.Parameters)
	if err != nil {
		return false, errors.Wrap(err, "acting in universe")
	}
	return quit, nil
}

// buy handles a buy action
func (t *textUI) buy(arg []interface{}) (bool, error) {
	if len(arg) != 2 {
		return false, errors.Errorf("buy passed wrong number of args, expected 2, got %v", len(arg))
	}
	amount, ok := arg[0].(int)
	if !ok {
		return false, errors.Errorf("buy passed non-int %v", arg[0])
	}
	item, ok := arg[1].(string)
	if !ok {
		return false, errors.Errorf("buy passed non-string %v", arg[1])
	}

	return false, t.u.Buy(amount, item)
}

// move handles a go action
func (t *textUI) move(arg []interface{}) (bool, error) {
	if len(arg) != 1 {
		return false, errors.Errorf("move passed wrong number of args, expected 1, got %v", len(arg))
	}
	destination, ok := arg[0].(string)
	if !ok {
		return false, errors.Errorf("move passed non-string %v", arg[0])
	}

	return false, t.u.Go(destination)
}

// help handles a help action
func (t *textUI) help(arg []interface{}) (bool, error) {
	switch len(arg) {
	case 0:
		actionList, err := t.u.Help()
		if err != nil {
			return false, errors.Wrap(err, "unable to get help list")
		}
		for _, actionType := range actionList {
			description := action.Describe(actionType)
			t.s.Out.Println(description.ShortUsage, "-", description.Description)
		}
	case 1:
		command, ok := arg[0].(string)
		if !ok {
			return false, errors.Errorf("help passed non-string %v", arg[1])
		}
		command = strings.ToLower(command)
		description, err := parseCommand(command)
		if err != nil {
			return false, errors.Wrapf(err, "help passed an unknown command %v", command)
		}
		t.s.Out.Println(description.Usage)
	default:
		return false, errors.Errorf("move passed wrong number of args, expected 0 or 1, got %v", len(arg))
	}
	return false, nil
}

// look handles a look action
func (t *textUI) look(arg []interface{}) (bool, error) {
	name, description, path, err := t.u.Look()
	if err != nil {
		return false, errors.Wrap(err, "look failed")
	}

	t.s.Out.Printf("You are at %s, %s.\n", name, description)
	t.s.Out.Println(path)
	return false, nil
}

// mine handles a mine action
func (t *textUI) mine(arg []interface{}) (bool, error) {
	return false, t.u.Mine()
}

// quit handles a quit command
func (t *textUI) quit(arg []interface{}) (bool, error) {
	return true, nil
}

// sell handles a sell action
func (t *textUI) sell(arg []interface{}) (bool, error) {
	if len(arg) != 2 {
		return false, errors.Errorf("sell passed wrong number of args, expected 2, got %v", len(arg))
	}
	amount, ok := arg[0].(int)
	if !ok {
		return false, errors.Errorf("sell passed non-int %v", arg[0])
	}
	item, ok := arg[1].(string)
	if !ok {
		return false, errors.Errorf("sell passed non-string %v", arg[1])
	}

	return false, t.u.Sell(amount, item)
}
