package textui

import (
	"corsairtext/action"
	"corsairtext/e"
	"strings"

	"github.com/pkg/errors"
)

func (t *textUI) call(request Request) (bool, error) {
	var actHandler = map[action.Type]func([]interface{}) (bool, error){
		action.TypeBuy:  t.buy,
		action.TypeGo:   t.move,
		action.TypeHelp: t.help,
		action.TypeLook: t.look,
		action.TypeDig:  t.dig,
		action.TypeQuit: t.quit,
		action.TypeSell: t.sell,
	}
	handler, ok := actHandler[request.Type]
	if !ok {
		return false, errors.Errorf("internal: no handler for %v", request.Type)
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
		return false, errors.Errorf("internal: buy passed wrong number of args, expected 2, got %v", len(arg))
	}
	amount, ok := arg[0].(int)
	if !ok {
		return false, errors.Errorf("internal: buy passed non-int %v", arg[0])
	}
	item, ok := arg[1].(string)
	if !ok {
		return false, errors.Errorf("internal: buy passed non-string %v", arg[1])
	}

	return false, t.u.Buy(amount, item)
}

// move handles a go action
func (t *textUI) move(arg []interface{}) (bool, error) {
	switch len(arg) {
	case 0:
		adjacency, err := t.u.GoList()
		if err != nil {
			return false, errors.Wrap(err, "internal: unable to get adjacency list")
		}
		for _, neighbor := range adjacency {
			t.s.Out.Println(neighbor.Name)
		}
	case 1:
		destination, ok := arg[0].(string)
		if !ok {
			return false, errors.Errorf("internal: move passed non-string %v", arg[0])
		}

		err := t.u.Go(destination)
		if err != nil {
			return false, err
		}

		t.look([]interface{}{nil})
	default:
		return false, errors.Errorf("internal: move passed wrong number of args, expected 0 or 1, got %v", len(arg))
	}
	return false, nil
}

// help handles a help action
func (t *textUI) help(arg []interface{}) (bool, error) {
	switch len(arg) {
	case 0:
		actionList, err := t.u.Help()
		if err != nil {
			return false, errors.Wrap(err, "internal: unable to get help list")
		}
		for _, actionType := range actionList {
			description := describe(actionType)
			usage := description.ShortUsage
			if usage == "" {
				usage = description.Usage
			}
			if usage == "" {
				continue
			}
			t.s.Out.Println(usage)
		}
	case 1:
		command, ok := arg[0].(string)
		if !ok {
			return false, errors.Errorf("internal: help passed non-string %v", arg[0])
		}
		command = strings.ToLower(command)
		description, err := parseCommand(command)
		if err != nil {
			return false, e.NewUnknownCommandError(command)
		}
		t.s.Out.Println(description.Usage)
	default:
		return false, errors.Errorf("internal: help passed wrong number of args, expected 0 or 1, got %v", len(arg))
	}
	return false, nil
}

// look handles a look action
func (t *textUI) look(arg []interface{}) (bool, error) {
	view, err := t.u.Look()
	if err != nil {
		return false, errors.Wrap(err, "internal: look failed")
	}

	t.s.Out.Println(strings.Join([]string{"You are at ", view.Name, ", ", view.Description, "."}, ""))
	t.s.Out.Println(view.Path)
	return false, nil
}

// dig handles a mining action
func (t *textUI) dig(arg []interface{}) (bool, error) {
	return false, t.u.Dig()
}

// quit handles a quit command
func (t *textUI) quit(arg []interface{}) (bool, error) {
	return true, nil
}

// sell handles a sell action
func (t *textUI) sell(arg []interface{}) (bool, error) {
	if len(arg) != 2 {
		return false, errors.Errorf("internal: sell passed wrong number of args, expected 2, got %v", len(arg))
	}
	amount, ok := arg[0].(int)
	if !ok {
		return false, errors.Errorf("internal: sell passed non-int %v", arg[0])
	}
	item, ok := arg[1].(string)
	if !ok {
		return false, errors.Errorf("internal: sell passed non-string %v", arg[1])
	}

	return false, t.u.Sell(amount, item)
}
