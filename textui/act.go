package textui

import (
	"corsairtext/action"

	"github.com/pkg/errors"
)

// act handles a user's command
func (t *textUI) act(command string, actionList action.List) (bool, error) {
	var actHandler = map[action.Type]func(interface{}) (bool, error){
		action.TypeLook: t.handleLook,
	}

	// parse the command from the command line
	request, err := t.parseAction(command, actionList)
	if err != nil {
		return false, errors.Wrap(err, "parsing action")
	}

	// send to universe
	response, err := t.u.Act(request)
	if err != nil {
		return false, errors.Wrap(err, "acting")
	}

	// handle the response
	handler, ok := actHandler[request.Type]
	if !ok {
		return false, errors.Errorf("no handler for %v", request.Type)
	}

	return handler(response)
}

// handleLook decodes the return from a look command
func (t *textUI) handleLook(rawResponse interface{}) (bool, error) {
	response, ok := rawResponse.(action.LookResponse)
	if !ok {
		return false, errors.New("look returned non-LookResponse")
	}

	t.s.Out.Printf("You are at %s, %s.\n", response.Name, response.Description)
	t.s.Out.Println(response.Path)
	return false, nil
}
