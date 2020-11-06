package textui

import (
	"corsairtext/action"
	"corsairtext/e"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Request is a parsed command
type Request struct {
	Type       action.Type
	Parameters []interface{}
}

// parse matches the input string with an action's regex
func (t *textUI) parse(input string) (Request, error) {
	var (
		rawWords      []string = strings.Split(input, " ")
		rawCommand    string   = strings.ToLower(rawWords[0])
		rawParameters []string = rawWords[1:]

		request           Request
		targetDescription *actionDescription
	)

	// find the command
	targetDescription, err := t.parseCommand(rawCommand)
	if err != nil {
		return request, err
	}
	request.Type = targetDescription.Type

	// validate the parameters
	request.Parameters, err = parseParameters(targetDescription, rawParameters)
	return request, err
}

// parseCommand matches a command to an action
func (t *textUI) parseCommand(rawCommand string) (*actionDescription, error) {
	commands := t.commandMatcher.Match(rawCommand)
	if len(commands) != 1 {
		return nil, e.NewUnknownCommandError(rawCommand)
	}
	command := commands[0]
	for _, description := range actionDescriptionTable {
		regexQuery := `\b` + description.NameRegex + `\b`
		match, err := regexp.MatchString(regexQuery, command)
		if err != nil {
			return nil, errors.Wrapf(err, "internal: bad match string, regex:%v command:%v", regexQuery, command)
		}
		if !match {
			continue
		}
		return &description, nil
	}
	return nil, e.NewUnknownCommandError(command)
}

// parseParameters loops through the parameters an parses them out
func parseParameters(targetDescription *actionDescription, rawParameters []string) ([]interface{}, error) {
	var parameters []interface{}
	count := 0
	for ; count < len(targetDescription.Parameters); count++ {
		parameterType := targetDescription.Parameters[count]
		regex, ok := parameterRegex[parameterType]
		if !ok {
			return nil, errors.Errorf("internal: unknown parameter type %v (index %v)", parameterType, count)
		}

		// Missing actual parameters, it only a problem it they are not optional
		if count >= len(rawParameters) {
			if parameterType != parameterTypeOptNumber && parameterType != parameterTypeOptAny {
				return nil, e.NewMissingParameterError(targetDescription.Type, len(targetDescription.Parameters), len(rawParameters))
			}
			break
		}

		// validate this parameter
		match, err := regexp.MatchString(regex, rawParameters[count])
		if err != nil {
			return nil, errors.Wrapf(err, "internal: malformed parameter #%v (%v) of type %v", count, rawParameters[count], parameterType)
		}
		if !match {
			return nil, e.NewBadParameterError(targetDescription.Type, rawParameters[count])
		}

		// convert this parameter
		switch parameterType {
		case parameterTypeNumber, parameterTypeOptNumber:
			value, err := strconv.Atoi(rawParameters[count])
			if err != nil {
				return parameters, errors.Errorf("internal: unable to convert parameter %v (%v) to a number", count, rawParameters[count])
			}
			parameters = append(parameters, value)
		case parameterTypeAny, parameterTypeOptAny:
			parameters = append(parameters, rawParameters[count])
		}
	}
	if count < len(rawParameters) {
		return nil, e.NewExtraParameterError(targetDescription.Type, len(targetDescription.Parameters), len(rawParameters))
	}

	return parameters, nil
}

// parameterRegex provides the proper regex for parameter types
var parameterRegex = map[parameterType]string{
	parameterTypeNone:      `\b`,
	parameterTypeNumber:    `\b\d+\b`,
	parameterTypeAny:       `\b.+\b`,
	parameterTypeOptNumber: `\b\d+\b`,
	parameterTypeOptAny:    `\b.+\b`,
}
