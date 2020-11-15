package commandprocessor

import (
	"corsairtext/e"
	"regexp"
	"strconv"
)

// parseCommand matches a string to a command
func (cp *commandProcessor) parseCommand(rawCommand string) (*commandDescription, error) {
	commands := cp.commandMatcher.Match(rawCommand)
	switch len(commands) {
	case 0:
		return nil, e.NewUnknownCommandError(rawCommand)
	default:
		return nil, e.NewAmbiguousCommandError(rawCommand, commands)
	case 1:
		description, ok := cp.descriptions[commands[0]]
		if !ok {
			return nil, e.NewUnknownCommandError(commands[0])
		}
		return description, nil
	}
}

func (cp *commandProcessor) parseNumber(text string) (int, bool) {
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
