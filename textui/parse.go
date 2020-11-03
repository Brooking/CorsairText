package textui

import (
	"corsairtext/action"
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

// parseAction matches the input string with an action's regex
func (t *textUI) parseAction(input string) (Request, error) {
	var (
		request           Request
		targetDescription *actionDescription

		rawWords      []string = strings.Split(input, " ")
		rawCommand    string   = strings.ToLower(rawWords[0])
		rawParameters []string = rawWords[1:]
	)

	// find the command
	targetDescription, err := parseCommand(rawCommand)
	if err != nil {
		return request, errors.Wrapf(err, "could not find %v", rawCommand)
	}
	request.Type = targetDescription.Type

	// validate the parameters
	var parameters []interface{}
	i := 0
	for ; i < len(targetDescription.Parameters); i++ {
		parameterType := targetDescription.Parameters[i]
		regex, ok := parameterRegex[parameterType]
		if !ok {
			return request, errors.Errorf("unknown parameter type %v (index %v)", parameterType, i)
		}

		if i >= len(rawParameters) {
			if parameterType != parameterTypeOptNumber && parameterType != parameterTypeOptAny {
				return request, errors.Errorf("missing parameters (expected %v, got %v)", len(targetDescription.Parameters), len(rawParameters))
			}
			break
		}

		match, err := regexp.MatchString(regex, rawParameters[i])
		if err != nil {
			return request, errors.Wrapf(err, "malformed parameter #%v (%v) of type %v", i, rawParameters[i], parameterType)
		}
		if !match {
			return request, errors.Errorf("malformed parameter %v (%v) of type %v", i, rawParameters[i], parameterType)
		}

		switch parameterType {
		case parameterTypeNumber, parameterTypeOptNumber:
			value, err := strconv.Atoi(rawParameters[i])
			if err != nil {
				return request, errors.Errorf("unable to convert parameter %v (%v) to a number", i, rawParameters[i])
			}
			parameters = append(parameters, value)
		case parameterTypeAny, parameterTypeOptAny:
			parameters = append(parameters, rawParameters[i])
		}
	}
	if i < len(rawParameters) {
		return request, errors.Errorf("too many parameters (expected %v, got %v)", len(targetDescription.Parameters), len(rawParameters))
	}

	request.Parameters = parameters
	return request, nil
}

// parseCommand matches a command to an action
func parseCommand(command string) (*actionDescription, error) {
	for _, description := range actionDescriptionTable {
		regexQuery := `\b` + description.NameRegex + `\b`
		match, err := regexp.MatchString(regexQuery, command)
		if err != nil {
			return nil, errors.Wrapf(err, "bad match string, regex:%v command:%v", regexQuery, command)
		}
		if !match {
			continue
		}
		return &description, nil
	}
	return nil, errors.Errorf("failed to match the command %v", command)
}

// parameterRegex provides the proper regex for parameter types
var parameterRegex = map[parameterType]string{
	parameterTypeNone:      `\b`,
	parameterTypeNumber:    `\b\d+\b`,
	parameterTypeAny:       `\b.+\b`,
	parameterTypeOptNumber: `\b\d+\b`,
	parameterTypeOptAny:    `\b.+\b`,
}
