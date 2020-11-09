package textui

import (
	"corsairtext/e"
	"regexp"
	"strconv"
	"strings"
)

// Request is a parsed command
// parse matches the input string with an action's regex
func (t *textUI) parse(input string) (interface{}, error) {
	var (
		rawWords      []string = strings.Split(input, " ")
		rawCommand    string   = strings.ToLower(rawWords[0])
		rawParameters []string = rawWords[1:]

		targetDescription *actionDescription
	)

	// find the command
	targetDescription, err := t.parseCommand(rawCommand)
	if err != nil {
		return nil, err
	}

	// validate the parameters
	return targetDescription.ParseParameters(t, rawParameters)
}

// parseCommand matches a command to an action
func (t *textUI) parseCommand(rawCommand string) (*actionDescription, error) {
	commands := t.commandMatcher.Match(rawCommand)
	switch len(commands) {
	case 0:
		return nil, e.NewUnknownCommandError(rawCommand)
	case 1:
		command := commands[0]
		for _, description := range actionDescriptionTable {
			// TODO move this into a map for fast lookup
			if command != description.Name {
				continue
			}
			return description, nil
		}
		return nil, e.NewUnknownCommandError(command)
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
