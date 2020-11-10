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
	return false, t.a.Buy(request.Amount, request.Item)
}

type goRequest struct {
	Destination string
}

// _go handles a go action
func (t *textUI) _go(request *goRequest) (bool, error) {
	switch {
	case request.Destination == "":
		adjacency := t.i.ListAdjacentLocations()
		for _, neighbor := range adjacency {
			t.s.Out.Println(neighbor)
		}
	default:
		err := t.a.Go(request.Destination)
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
		for _, command := range t.i.ListLocalCommands() {
			description, ok := actionDescriptionMap[command]
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
		commands := t.commandMatcher.Match(request.Command)
		switch len(commands) {
		case 0:
			return false, e.NewUnknownCommandError(request.Command)
		case 1:
			description, ok := actionDescriptionMap[request.Command]
			if !ok {
				return false, e.NewUnknownCommandError(request.Command)
			}
			t.s.Out.Println(description.Usage)
			return false, nil
		default:
			return false, e.NewUnknownCommandError(request.Command)
		}
	}
	return false, nil
}

type lookRequest struct{}

// look handles a look action
func (t *textUI) look(request *lookRequest) (bool, error) {
	location := t.i.LocalLocation()
	t.s.Out.Println(strings.Join([]string{"You are at ", location.Name, ", ", location.Description, "."}, ""))

	var path string
	for _, spot := range location.Path {
		path = path + spot + "/"
	}
	t.s.Out.Println(path)
	return false, nil
}

type digRequest struct{}

// dig handles a mining action
func (t *textUI) dig(request *digRequest) (bool, error) {
	return false, t.a.Dig()
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
	return false, t.a.Sell(request.Amount, request.Item)
}
