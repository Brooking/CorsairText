package commandprocessor

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
				assert.Equal(t, "help", description.ShortName)
			},
		},
		{
			name:  "matches hel",
			input: "hel",
			assert: func(description *commandDescription, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "help", description.ShortName)
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
				assert.Equal(t, "help", description.ShortName)
			},
		},
		{
			name:  "matches Go without parameters",
			input: "G",
			assert: func(description *commandDescription, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "go", description.ShortName)
			},
		},
		{
			name:  "matches Sell",
			input: "Sell",
			assert: func(description *commandDescription, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "sell", description.ShortName)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			cp := NewCommandProcessor(nil, nil, nil)
			realcp, ok := cp.(*commandProcessor)
			assert.True(t, ok)

			// act
			commandDescription, err := realcp.parseCommand(testCase.input)

			// assert
			testCase.assert(commandDescription, err)

		})
	}
}
