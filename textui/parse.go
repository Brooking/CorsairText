package textui

import (
	"corsairtext/e"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
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
	return targetDescription.ParseParameters(rawParameters)
}

// parseCommand matches a command to an action
func (t *textUI) parseCommand(rawCommand string) (*actionDescription, error) {
	commands := t.commandMatcher.Match(rawCommand)
	switch len(commands) {
	case 0:
		return nil, e.NewUnknownCommandError(rawCommand)
	case 1:
		command, ok := commands[0].Context.(*actionDescription)
		if !ok {
			return nil, errors.Errorf("something other than a command decision came back from command matcher %T", commands[0].Context)
		}
		return command, nil
	default:
		var commandList []string
		for _, entry := range commands {
			commandList = append(commandList, entry.Word)
		}
		return nil, e.NewAmbiguousCommandError(rawCommand, commandList)
	}
}

func parseNumber(text string) (int, bool) {
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
