package textui

import (
	"corsairtext/e"
	"strings"

	"github.com/pkg/errors"
)

func (t *textUI) call(request interface{}) (bool, error) {
	switch r := request.(type) {
	case *buyRequest:
		return t.buy(r)
	case *digRequest:
		return t.dig(r)
	case *goRequest:
		return t._go(r)
	case *helpRequest:
		return t.help(r)
	case *lookRequest:
		return t.look(r)
	case *quitRequest:
		return t.quit(r)
	case *sellRequest:
		return t.sell(r)
	default:
		return false, errors.Errorf("internal: no call handler for %T", r)
	}
}

type buyRequest struct {
	Amount int
	Item   string
}

// buy handles a buy action
func (t *textUI) buy(request *buyRequest) (bool, error) {
	return false, t.u.Buy(request.Amount, request.Item)
}

type goRequest struct {
	Destination string
}

// _go handles a go action
func (t *textUI) _go(request *goRequest) (bool, error) {
	switch {
	case request.Destination == "":
		adjacency, err := t.u.GoList()
		if err != nil {
			return false, errors.Wrap(err, "internal: unable to get adjacency list")
		}
		for _, neighbor := range adjacency {
			t.s.Out.Println(neighbor.Name)
		}
	default:
		err := t.u.Go(request.Destination)
		if err != nil {
			return false, err
		}

		t.look(&lookRequest{})
	}
	return false, nil
}

type helpRequest struct {
	Command string
}

// help handles a help action
func (t *textUI) help(request *helpRequest) (bool, error) {
	switch {
	case request.Command == "":
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
	default:
		command := strings.ToLower(request.Command)
		description, err := t.parseCommand(command)
		if err != nil {
			return false, e.NewUnknownCommandError(command)
		}
		t.s.Out.Println(description.Usage)
	}
	return false, nil
}

type lookRequest struct{}

// look handles a look action
func (t *textUI) look(request *lookRequest) (bool, error) {
	view, err := t.u.Look()
	if err != nil {
		return false, errors.Wrap(err, "internal: look failed")
	}

	t.s.Out.Println(strings.Join([]string{"You are at ", view.Name, ", ", view.Description, "."}, ""))

	var path string
	for _, spot := range view.Path {
		path = path + spot + "/"
	}
	t.s.Out.Println(path)
	return false, nil
}

type digRequest struct{}

// dig handles a mining action
func (t *textUI) dig(request *digRequest) (bool, error) {
	return false, t.u.Dig()
}

type quitRequest struct{}

// quit handles a quit command
func (t *textUI) quit(request *quitRequest) (bool, error) {
	return true, nil
}

type sellRequest struct {
	Amount int
	Item   string
}

// sell handles a sell action
func (t *textUI) sell(request *sellRequest) (bool, error) {
	return false, t.u.Sell(request.Amount, request.Item)
}
