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
		request                  Request
		matchedActionDescription *action.Description

		words      []string = strings.Split(input, " ")
		command    string   = strings.ToLower(words[0])
		parameters []string = words[1:]
	)

	// find the command
	matchedActionDescription, err := parseCommand(command)
	if err != nil {
		return request, errors.Wrapf(err, "could not find %v", command)
	}
	request.Type = matchedActionDescription.Type

	// validate the parameters
	if len(parameters) > len(matchedActionDescription.Parameters) {
		return request, errors.Errorf("missing parameters (expected %v, got %v)", len(matchedActionDescription.Parameters), len(parameters))
	}

	if len(parameters) < len(matchedActionDescription.Parameters) {
		return request, errors.Errorf("too many parameters (expected %v, got %v)", len(matchedActionDescription.Parameters), len(parameters))
	}

	for i := 0; i < len(matchedActionDescription.Parameters); i++ {
		parameterType := matchedActionDescription.Parameters[i]
		regex, ok := parameterRegex[parameterType]
		if !ok {
			return request, errors.Errorf("unknown parameter type %v (index %v)", parameterType, i)
		}

		match, err := regexp.MatchString(regex, parameters[i])
		if err != nil {
			return request, errors.Wrapf(err, "malformed parameter #%v (%v) of type %v", i, parameters[i], parameterType)
		}
		if !match {
			return request, errors.Errorf("malformed parameter %v (%v) of type %v", i, parameters[i], parameterType)
		}

		switch parameterType {
		case action.ParameterTypeNumber:
			value, err := strconv.Atoi(parameters[i])
			if err != nil {
				return request, errors.Errorf("unable to convert parameter %v (%v) to a number", i, parameters[i])
			}
			request.Parameters = append(request.Parameters, value)
		case action.ParameterTypeAny:
			request.Parameters = append(request.Parameters, parameters[i])
		}
	}

	return request, nil
}

// parseCommand matches a command to an action
func parseCommand(command string) (*action.Description, error) {
	for _, description := range action.DescriptionTable {
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
var parameterRegex = map[action.ParameterType]string{
	action.ParameterTypeNone:   `\b`,
	action.ParameterTypeNumber: `\b\d+\b`,
	action.ParameterTypeAny:    `\b.+\b`,
}
