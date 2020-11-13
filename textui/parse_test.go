package textui

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestParseCommandLine(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		assert func(*commandDescription, error)
	}{
		{
			name:  "empty",
			input: "",
			assert: func(description *commandDescription, err error) {
				assert.Error(t, err)
			},
		},
		{
			name:  "matches h",
			input: "h",
			assert: func(description *commandDescription, err error) {
				assert.NoError(t, err)
				assert.Equal(t, commandDescriptionMap[CommandHelp], description)
			},
		},
		{
			name:  "matches hel",
			input: "hel",
			assert: func(description *commandDescription, err error) {
				assert.NoError(t, err)
				assert.Equal(t, commandDescriptionMap[CommandHelp], description)
			},
		},
		{
			name:  "fails helps",
			input: "helps",
			assert: func(description *commandDescription, err error) {
				assert.Error(t, err)
			},
		},
		{
			name:  "matches Help",
			input: "Help",
			assert: func(description *commandDescription, err error) {
				assert.NoError(t, err)
				assert.Equal(t, commandDescriptionMap[CommandHelp], description)
			},
		},
		{
			name:  "matches Go without parameters",
			input: "G",
			assert: func(description *commandDescription, err error) {
				assert.NoError(t, err)
				assert.Equal(t, commandDescriptionMap[CommandGo], description)
			},
		},
		{
			name:  "matches Sell",
			input: "Sell",
			assert: func(description *commandDescription, err error) {
				assert.NoError(t, err)
				assert.Equal(t, commandDescriptionMap[CommandSell], description)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			textui := &textUI{
				commandMatcher: NewCommandMatcher(),
			}

			// act
			command, err := textui.parseCommand(testCase.input)

			// assert
			testCase.assert(command, err)

		})
	}
}
