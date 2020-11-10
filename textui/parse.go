package textui

import (
	"corsairtext/e"
	"regexp"
	"strconv"
	"strings"
)

// parse matches the input string with a command's regex
func (t *textUI) parse(input string) (interface{}, error) {
	var (
		words       []string = strings.Split(input, " ")
		command     string   = strings.ToLower(words[0])
		parameters  []string = words[1:]
		description *commandDescription
		err         error
	)

	// find the command
	description, err = t.parseCommand(command)
	if err != nil {
		return nil, err
	}

	// validate the parameters
	return description.ParseParameters(t, parameters)
}

// parseCommand matches a string to a command
func (t *textUI) parseCommand(rawCommand string) (*commandDescription, error) {
	commands := t.commandMatcher.Match(rawCommand)
	switch len(commands) {
	case 0:
		return nil, e.NewUnknownCommandError(rawCommand)
	case 1:
		description, ok := commandDescriptionMap[commands[0]]
		if !ok {
			return nil, e.NewUnknownCommandError(commands[0])
		}
		return description, nil
	default:
		return nil, e.NewAmbiguousCommandError(rawCommand, commands)
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
