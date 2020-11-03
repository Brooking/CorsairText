package textui

import (
	"corsairtext/action"

	"github.com/pkg/errors"
)

// act handles a user's command
func (t *textUI) act(command string, actionList action.List) (bool, error) {
	var actHandler = map[action.Type]func(interface{}) (bool, error){
		action.TypeHelp: t.handleHelp,
		action.TypeLook: t.handleLook,
		action.TypeQuit: t.handleQuit,
	}

	// parse the command from the command line
	request, err := t.parseAction(command, actionList)
	if err != nil {
		return false, errors.Wrap(err, "parsing action")
	}

	// send to universe
	response, err := t.u.Act(request)
	if err != nil {
		return false, errors.Wrap(err, "acting in universe")
	}

	// handle the response
	handler, ok := actHandler[request.Type]
	if !ok {
		return false, errors.Errorf("no handler for %v", request.Type)
	}

	return handler(response)
}

// handleHelp decodes the return from a help command
func (t *textUI) handleHelp(rawResponse interface{}) (bool, error) {
	response, ok := rawResponse.([]action.Description)
	if !ok {
		return false, errors.New("help returned non-HelpResponse")
	}

	for _, description := range response {
		t.s.Out.Println(description.ShortUsage, "-", description.Description)
	}
	return false, nil
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

// handleQuit decodes the return from a quit command
func (t *textUI) handleQuit(rawResponse interface{}) (bool, error) {
	return true, nil
}
