package textui

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestParseAction(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		assert func(interface{}, error)
	}{
		{
			name:  "empty",
			input: "",
			assert: func(request interface{}, err error) {
				assert.Error(t, err)
			},
		},
		{
			name:  "matches h",
			input: "h",
			assert: func(request interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &helpRequest{}, request)
			},
		},
		{
			name:  "matches hel",
			input: "hel",
			assert: func(request interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &helpRequest{}, request)
			},
		},
		{
			name:  "fails helps",
			input: "helps",
			assert: func(request interface{}, err error) {
				assert.Error(t, err)
			},
		},
		{
			name:  "matches Help",
			input: "Help",
			assert: func(request interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &helpRequest{}, request)
			},
		},
		{
			name:  "matches Help with one parameter",
			input: "Help Buy",
			assert: func(request interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &helpRequest{}, request)
				r := request.(*helpRequest)
				assert.Equal(t, "Buy", r.Command)
			},
		},
		{
			name:  "fails Help with two parameters",
			input: "Help look go",
			assert: func(request interface{}, err error) {
				assert.Error(t, err)
			},
		},
		{
			name:  "matches Go without parameters",
			input: "G",
			assert: func(request interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &goRequest{}, request)
				r := request.(*goRequest)
				assert.Equal(t, "", r.Destination)
			},
		},
		{
			name:  "matches Go with 1 parameter",
			input: "G moon",
			assert: func(request interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &goRequest{}, request)
				r := request.(*goRequest)
				assert.Equal(t, "moon", r.Destination)
			},
		},
		{
			name:  "fails Go 2 parameters",
			input: "Go to mars",
			assert: func(request interface{}, err error) {
				assert.Error(t, err)
			},
		},
		{
			name:  "matches Sell with good parameters",
			input: "Sell 14 ore",
			assert: func(request interface{}, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &sellRequest{}, request)
				r := request.(*sellRequest)
				assert.Equal(t, 14, r.Amount)
				assert.Equal(t, "ore", r.Item)
			},
		},
		{
			name:  "fails Sell with bad parameter",
			input: "Sell fifty computers",
			assert: func(request interface{}, err error) {
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
			request, err := textui.parse(testCase.input)

			// assert
			testCase.assert(request, err)

		})
	}
}
