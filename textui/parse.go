package textui

import (
	"corsairtext/action"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// parseAction matches the input string with an action's regex
func (*textUI) parseAction(input string, actions action.List) (action.Type, error) {
	var matchedAction *action.Description
	words := strings.Split(input, " ")

	// find the command
	command := strings.ToLower(words[0])
	for _, description := range actions.Descriptions() {
		match, err := regexp.MatchString("\\b"+description.Regex[0]+"\\b", command)
		if err != nil {
			continue
		}
		if !match {
			continue
		}
		matchedAction = &description
		break
	}

	if matchedAction == nil {
		return action.TypeNone, errors.New("failed to match the command")
	}

	// validate the parameters
	if len(words) > len(matchedAction.Regex) {
		return matchedAction.Type, errors.New("missing parameters")
	}

	if len(words) < len(matchedAction.Regex) {
		return matchedAction.Type, errors.New("too many parameters")
	}

	for i := 1; i < len(matchedAction.Regex); i++ {
		match, err := regexp.MatchString("\\b"+matchedAction.Regex[i]+"\\b", words[i])
		if err != nil {
			return matchedAction.Type, errors.Wrapf(err, "malformed parameter #%v", i)
		}
		if !match {
			return matchedAction.Type, errors.Errorf("malformed parameter %v", i)
		}
	}

	return matchedAction.Type, nil
}
