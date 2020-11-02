package textui

import (
	"corsairtext/action"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// parseAction matches the input string with an action's regex
func (*textUI) parseAction(input string, actions action.List) (action.Request, error) {
	var (
		request                  action.Request
		matchedActionDescription *action.Description

		words      []string = strings.Split(input, " ")
		command    string   = strings.ToLower(words[0])
		parameters []string = words[1:]
	)

	// find the command
	for _, description := range actions.Descriptions() {
		regexQuery := `\b` + description.NameRegex + `\b`
		match, err := regexp.MatchString(regexQuery, command)
		if err != nil {
			return request, errors.Wrapf(err, "bad match string, regex:%v command:%v", regexQuery, command)
		}
		if !match {
			continue
		}
		matchedActionDescription = &description
		break
	}

	if matchedActionDescription == nil {
		return request, errors.Errorf("failed to match the command %v", command)
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

// parameterRegex provides the proper regex for parameter types
var parameterRegex = map[action.ParameterType]string{
	action.ParameterTypeNone:   `\b`,
	action.ParameterTypeNumber: `\b\d+\b`,
	action.ParameterTypeAny:    `\b.+\b`,
}
