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
		assert func(interface{}, error)
	}{
		{
			name:  "empty",
			input: "",
			assert: func(command interface{}, err error) {
				assert.Error(t, err)
			},
		},
		{
			name:  "matches h",
			input: "h",
			assert: func(command interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &helpCommand{}, command)
			},
		},
		{
			name:  "matches hel",
			input: "hel",
			assert: func(command interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &helpCommand{}, command)
			},
		},
		{
			name:  "fails helps",
			input: "helps",
			assert: func(command interface{}, err error) {
				assert.Error(t, err)
			},
		},
		{
			name:  "matches Help",
			input: "Help",
			assert: func(command interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &helpCommand{}, command)
			},
		},
		{
			name:  "matches Help with one parameter",
			input: "Help Buy",
			assert: func(command interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &helpCommand{}, command)
				r := command.(*helpCommand)
				assert.Equal(t, "Buy", r.Command)
			},
		},
		{
			name:  "fails Help with two parameters",
			input: "Help look go",
			assert: func(command interface{}, err error) {
				assert.Error(t, err)
			},
		},
		{
			name:  "matches Go without parameters",
			input: "G",
			assert: func(command interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &goCommand{}, command)
				r := command.(*goCommand)
				assert.Equal(t, "", r.Destination)
			},
		},
		{
			name:  "matches Go with 1 parameter",
			input: "G moon",
			assert: func(command interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &goCommand{}, command)
				r := command.(*goCommand)
				assert.Equal(t, "moon", r.Destination)
			},
		},
		{
			name:  "fails Go 2 parameters",
			input: "Go to mars",
			assert: func(command interface{}, err error) {
				assert.Error(t, err)
			},
		},
		{
			name:  "matches Sell with good parameters",
			input: "Sell 14 ore",
			assert: func(command interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &sellCommand{}, command)
				r := command.(*sellCommand)
				assert.Equal(t, 14, r.Amount)
				assert.Equal(t, "ore", r.Item)
			},
		},
		{
			name:  "fails Sell with bad parameter",
			input: "Sell fifty computers",
			assert: func(command interface{}, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			textui := &textUI{
				commandMatcher: MakeCommandMatcher(),
			}

			// act
			command, err := textui.parse(testCase.input)

			// assert
			testCase.assert(command, err)

		})
	}
}
