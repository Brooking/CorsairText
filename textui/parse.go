package textui

import (
	"corsairtext/e"
	"regexp"
	"strconv"
	"strings"
)

// act handles a user's command
func (t *textUI) act(commandString string) error {
	var (
		words       []string = strings.Split(commandString, " ")
		command     string   = strings.ToLower(words[0])
		parameters  []string = words[1:]
		description *commandDescription
		err         error
	)

	// find the command
	description, err = t.parseCommand(command)
	if err != nil {
		return err
	}

	// handle the command
	return description.Handler(t, parameters)
}

// parseCommand matches a string to a command
func (t *textUI) parseCommand(rawCommand string) (*commandDescription, error) {
	commands := t.commandMatcher.Match(rawCommand)
	switch len(commands) {
	case 0:
		return nil, e.NewUnknownCommandError(rawCommand)
	default:
		return nil, e.NewAmbiguousCommandError(rawCommand, commands)
	case 1:
		description, ok := commandDescriptionMap[commands[0]]
		if !ok {
			return nil, e.NewUnknownCommandError(commands[0])
		}
		return description, nil
	}
}

func (t *textUI) parseNumber(text string) (int, bool) {
	match, err := regexp.MatchString(`\b\d+\b`, text)
	if err != nil {
		return 0, false
	}
	if !match {
		return 0, false
	}
	value, err := strconv.Atoi(text)
	if err != nil {
		return 0, false
	}
	return value, true
}
