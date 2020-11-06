package textui

import (
	"corsairtext/action"
	"corsairtext/match/mockmatch"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestParseAction(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		assert func(Request, error)
	}{
		{
			name:  "empty",
			input: "",
			assert: func(request Request, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeNone, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "matches h",
			input: "h",
			assert: func(request Request, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeHelp, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "fails he",
			input: "he",
			assert: func(request Request, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeNone, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "matches Help",
			input: "Help",
			assert: func(request Request, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeHelp, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "matches Help with one parameter",
			input: "Help Buy",
			assert: func(request Request, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeHelp, request.Type, "action type")
				assert.Equal(t, 1, len(request.Parameters), "# parameters")
				assert.Equal(t, "Buy", request.Parameters[0], "parameter 0")
			},
		},
		{
			name:  "fails Help with two parameters",
			input: "Help look go",
			assert: func(request Request, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeHelp, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "fails Go without parameters",
			input: "G",
			assert: func(request Request, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeGo, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "matches Go with 1 parameter",
			input: "G moon",
			assert: func(request Request, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeGo, request.Type, "action type")
				assert.Equal(t, 1, len(request.Parameters), "# parameters")
				assert.Equal(t, "moon", request.Parameters[0], "parameter 0")
			},
		},
		{
			name:  "fails Go 2 parameters",
			input: "Go to mars",
			assert: func(request Request, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeGo, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
		{
			name:  "matches Sell with good parameters",
			input: "Sell 14 ore",
			assert: func(request Request, err error) {
				assert.NoError(t, err)
				assert.Equal(t, action.TypeSell, request.Type, "action type")
				assert.Equal(t, 2, len(request.Parameters), "# parameters")
				assert.Equal(t, 14, request.Parameters[0], "parameter 0")
				assert.Equal(t, "ore", request.Parameters[1], "parameter 1")
			},
		},
		{
			name:  "fails Sell with bad parameter",
			input: "Sell fifty computers",
			assert: func(request Request, err error) {
				assert.Error(t, err)
				assert.Equal(t, action.TypeSell, request.Type, "action type")
				assert.Equal(t, 0, len(request.Parameters), "# parameters")
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			matcherMock := mockmatch.NewMockMatcher(ctrl)
			matcherMock.EXPECT().Ingest(gomock.Any()).AnyTimes()
			matcherMock.EXPECT().
				Match(gomock.Any()).
				DoAndReturn(func(s string) string {
					return s
				}).
				AnyTimes()
			textui := &textUI{
				commandMatcher: matcherMock,
			}

			// act
			request, err := textui.parse(testCase.input)

			// assert
			testCase.assert(request, err)

		})
	}
}
