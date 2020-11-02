package textui

import (
	"corsairtext/action"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAction(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		assert func(action.Type, error)
	}{
		{
			name:  "empty",
			input: "",
			assert: func(actionType action.Type, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeNone, actionType, "action type")
			},
		},
		{
			name:  "matches h",
			input: "h",
			assert: func(actionType action.Type, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeHelp, actionType, "action type")
			},
		},
		{
			name:  "fails he",
			input: "he",
			assert: func(actionType action.Type, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeNone, actionType, "action type")
			},
		},
		{
			name:  "matches Help",
			input: "Help",
			assert: func(actionType action.Type, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeHelp, actionType, "action type")
			},
		},
		{
			name:  "fails Help with parameters",
			input: "Help me",
			assert: func(actionType action.Type, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeHelp, actionType, "action type")
			},
		},
		{
			name:  "fails Go without parameters",
			input: "G",
			assert: func(actionType action.Type, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeGo, actionType, "action type")
			},
		},
		{
			name:  "matches Go with 1 parameter",
			input: "G moon",
			assert: func(actionType action.Type, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeGo, actionType, "action type")
			},
		},
		{
			name:  "fails Go 2 parameters",
			input: "Go to mars",
			assert: func(actionType action.Type, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeGo, actionType, "action type")
			},
		},
		{
			name:  "matches Sell with good parameters",
			input: "Sell 14 ore",
			assert: func(actionType action.Type, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeSell, actionType, "action type")
			},
		},
		{
			name:  "fails Sell with bad parameter",
			input: "Sell fifty computers",
			assert: func(actionType action.Type, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeSell, actionType, "action type")
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			textui := &textUI{}
			actionList := action.List{
				action.TypeBuy,
				action.TypeGo,
				action.TypeHelp,
				action.TypeLook,
				action.TypeMine,
				action.TypeSell,
			}

			// act
			actionType, err := textui.parseAction(testCase.input, actionList)

			// assert
			testCase.assert(actionType, err)

		})
	}
}
