package textui

import (
	"corsairtext/action"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestParseAction(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		actionList action.List
		assert     func(action.Type, error)
	}{
		{
			name:  "empty",
			input: "",
			actionList: []action.Type{
				action.TypeBuy,
				action.TypeGo,
				action.TypeHelp,
				action.TypeLook,
				action.TypeMine,
				action.TypeSell,
			},
			assert: func(actionType action.Type, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeNone, actionType, "action type")
			},
		},
		{
			name:  "matches h",
			input: "h",
			actionList: []action.Type{
				action.TypeHelp,
			},
			assert: func(actionType action.Type, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeHelp, actionType, "action type")
			},
		},
		{
			name:  "fails he",
			input: "he",
			actionList: []action.Type{
				action.TypeHelp,
			},
			assert: func(actionType action.Type, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeNone, actionType, "action type")
			},
		},
		{
			name:  "matches Help",
			input: "Help",
			actionList: []action.Type{
				action.TypeHelp,
			},
			assert: func(actionType action.Type, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeHelp, actionType, "action type")
			},
		},
		{
			name:  "fails Help with parameters",
			input: "Help me",
			actionList: []action.Type{
				action.TypeHelp,
			},
			assert: func(actionType action.Type, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeHelp, actionType, "action type")
			},
		},
		{
			name:  "fails Go without parameters",
			input: "G",
			actionList: []action.Type{
				action.TypeGo,
			},
			assert: func(actionType action.Type, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeGo, actionType, "action type")
			},
		},
		{
			name:  "matches Go with 1 parameter",
			input: "G moon",
			actionList: []action.Type{
				action.TypeGo,
			},
			assert: func(actionType action.Type, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeGo, actionType, "action type")
			},
		},
		{
			name:  "fails Go 2 parameters",
			input: "Go to mars",
			actionList: []action.Type{
				action.TypeGo,
			},
			assert: func(actionType action.Type, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeGo, actionType, "action type")
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			textui := &textUI{}

			// act
			actionType, err := textui.parseAction(testCase.input, testCase.actionList)

			// assert
			testCase.assert(actionType, err)

		})
	}
}
