package textui

import (
	"errors"
	"testing"

	"corsairtext/action"
	"corsairtext/support"
	"corsairtext/support/screenprinter/mockscreenprinter"
	"corsairtext/universe/mockuniverse"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	testCases := []struct {
		name    string
		request Request
		assert  func(bool, error)
	}{
		{
			name: "success quit",
			request: Request{
				Type: action.TypeQuit,
			},
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
				assert.Equal(t, true, quit)
			},
		},
		{
			name: "fail bad type",
			request: Request{
				Type: 1000,
			},
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			textui := &textUI{}

			// act
			quit, err := textui.call(testCase.request)

			// assert
			testCase.assert(quit, err)
		})
	}
}

func TestCallBuy(t *testing.T) {
	testCases := []struct {
		name      string
		request   Request
		buyAmount int
		buyItem   string
		buyReturn error
		buyCalls  int
		assert    func(bool, error)
	}{
		{
			name: "buy success",
			request: Request{
				Type:       action.TypeBuy,
				Parameters: []interface{}{3, "computers"},
			},
			buyAmount: 3,
			buyItem:   "computers",
			buyCalls:  1,
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "buy failed missing params",
			request: Request{
				Type:       action.TypeBuy,
				Parameters: []interface{}{3},
			},
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "buy failed bad first param",
			request: Request{
				Type:       action.TypeBuy,
				Parameters: []interface{}{"three", "computers"},
			},
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "buy failed bad second param",
			request: Request{
				Type:       action.TypeBuy,
				Parameters: []interface{}{3, nil},
			},
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "buy call fail",
			request: Request{
				Type:       action.TypeBuy,
				Parameters: []interface{}{3, "computers"},
			},
			buyAmount: 3,
			buyItem:   "computers",
			buyReturn: errors.New("some buy error"),
			buyCalls:  1,
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			universeMock := mockuniverse.NewMockAction(ctrl)
			universeMock.EXPECT().
				Buy(testCase.buyAmount, testCase.buyItem).
				Return(testCase.buyReturn).
				Times(testCase.buyCalls)
			textui := &textUI{
				u: universeMock,
			}

			// act
			quit, err := textui.call(testCase.request)

			// assert
			testCase.assert(quit, err)
		})
	}
}

func TestCallGo(t *testing.T) {
	testCases := []struct {
		name          string
		request       Request
		goDestination string
		goReturn      error
		goCalls       int
		assert        func(bool, error)
	}{
		{
			name: "go success",
			request: Request{
				Type:       action.TypeGo,
				Parameters: []interface{}{"mars"},
			},
			goDestination: "mars",
			goCalls:       1,
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "go failed missing param",
			request: Request{
				Type:       action.TypeGo,
				Parameters: []interface{}{},
			},
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "go failed bad param",
			request: Request{
				Type:       action.TypeGo,
				Parameters: []interface{}{nil},
			},
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "go call fail",
			request: Request{
				Type:       action.TypeGo,
				Parameters: []interface{}{"mars"},
			},
			goDestination: "mars",
			goReturn:      errors.New("some go error"),
			goCalls:       1,
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			universeMock := mockuniverse.NewMockAction(ctrl)
			universeMock.EXPECT().
				Go(testCase.goDestination).
				Return(testCase.goReturn).
				Times(testCase.goCalls)
			textui := &textUI{
				u: universeMock,
			}

			// act
			quit, err := textui.call(testCase.request)

			// assert
			testCase.assert(quit, err)
		})
	}
}

func TestCallHelp(t *testing.T) {
	testCases := []struct {
		name       string
		request    Request
		helpReturn []action.Type
		helpError  error
		helpCalls  int
		outInput   string
		outCalls   int
		assert     func(bool, error)
	}{
		{
			name: "success 0 params (Go)",
			request: Request{
				Type: action.TypeHelp,
			},
			helpReturn: []action.Type{
				action.TypeGo,
			},
			helpCalls: 1,
			outInput:  "(G)o - Travel",
			outCalls:  1,
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "success 0 params (Look)",
			request: Request{
				Type: action.TypeHelp,
			},
			helpReturn: []action.Type{
				action.TypeLook,
			},
			helpCalls: 1,
			outInput:  "(L)ook - Look around",
			outCalls:  1,
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "help call fail",
			request: Request{
				Type: action.TypeHelp,
			},
			helpError: errors.New("some go error"),
			helpCalls: 1,
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "success 1 param",
			request: Request{
				Type:       action.TypeHelp,
				Parameters: []interface{}{"Go"},
			},
			outInput: "(G)o <destination> - Travel to destination",
			outCalls: 1,
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "fail 1 nil param",
			request: Request{
				Type:       action.TypeHelp,
				Parameters: []interface{}{nil},
			},
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "fail 1 unknown param",
			request: Request{
				Type:       action.TypeHelp,
				Parameters: []interface{}{"DoAFlip"},
			},
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "fail 2 params",
			request: Request{
				Type:       action.TypeHelp,
				Parameters: []interface{}{"Go", "Travel"},
			},
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			universeMock := mockuniverse.NewMockAction(ctrl)
			universeMock.EXPECT().
				Help().
				Return(testCase.helpReturn, testCase.helpError).
				Times(testCase.helpCalls)
			outMock := mockscreenprinter.NewMockScreenPrinter(ctrl)
			outMock.EXPECT().
				Println(testCase.outInput).
				Return(0, nil).
				Times(testCase.outCalls)
			support := support.Support{
				Out: outMock,
			}
			textui := &textUI{
				s: support,
				u: universeMock,
			}

			// act
			quit, err := textui.call(testCase.request)

			// assert
			testCase.assert(quit, err)
		})
	}
}
