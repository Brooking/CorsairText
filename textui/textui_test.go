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
		assert func(action.Request, error)
	}{
		{
			name:  "empty",
			input: "",
			assert: func(request action.Request, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeNone, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "matches h",
			input: "h",
			assert: func(request action.Request, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeHelp, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "fails he",
			input: "he",
			assert: func(request action.Request, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeNone, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "matches Help",
			input: "Help",
			assert: func(request action.Request, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeHelp, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "fails Help with parameters",
			input: "Help me",
			assert: func(request action.Request, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeHelp, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "fails Go without parameters",
			input: "G",
			assert: func(request action.Request, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeGo, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "matches Go with 1 parameter",
			input: "G moon",
			assert: func(request action.Request, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeGo, request.Type, "action type")
				assert.Equal(t, 1, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "fails Go 2 parameters",
			input: "Go to mars",
			assert: func(request action.Request, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeGo, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "matches Sell with good parameters",
			input: "Sell 14 ore",
			assert: func(request action.Request, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeSell, request.Type, "action type")
				assert.Equal(t, 2, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "fails Sell with bad parameter",
			input: "Sell fifty computers",
			assert: func(request action.Request, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeSell, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
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
			request, err := textui.parseAction(testCase.input, actionList)

			// assert
			testCase.assert(request, err)

		})
	}
}
