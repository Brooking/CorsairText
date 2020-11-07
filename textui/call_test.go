package textui

import (
	"errors"
	"testing"

	"corsairtext/action"
	"corsairtext/e"
	"corsairtext/match/mockmatch"
	"corsairtext/support"
	"corsairtext/support/screenprinter/mockscreenprinter"
	"corsairtext/universe"
	"corsairtext/universe/mockuniverse"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	testCases := []struct {
		name    string
		request interface{}
		assert  func(bool, error)
	}{
		{
			name:    "success quit",
			request: &quitRequest{},
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
				assert.Equal(t, true, quit)
			},
		},
		{
			name:    "fail bad struct",
			request: e.AmbiguousCommandError{},
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name:    "fail nil struct",
			request: nil,
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
		request   interface{}
		buyAmount int
		buyItem   string
		buyReturn error
		buyCalls  int
		assert    func(bool, error)
	}{
		{
			name: "buy success",
			request: &buyRequest{
				Amount: 3,
				Item:   "computers",
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
			name: "buy call fail",
			request: &buyRequest{
				Amount: 3,
				Item:   "computers",
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
			matcherMock := mockmatch.NewMockMatcher(ctrl)
			matcherMock.EXPECT().Ingest(gomock.Any()).AnyTimes()
			matcherMock.EXPECT().
				Match(gomock.Any()).
				DoAndReturn(func(s string) []string {
					return []string{s}
				}).
				AnyTimes()
			textui := &textUI{
				u:              universeMock,
				commandMatcher: matcherMock,
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
		request       interface{}
		goDestination string
		goReturn      error
		goCalls       int
		lookCalls     int
		assert        func(bool, error)
	}{
		{
			name: "go success with dest",
			request: &goRequest{
				Destination: "mars",
			},
			goDestination: "mars",
			goCalls:       1,
			lookCalls:     1,
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "go call fail",
			request: &goRequest{
				Destination: "mars",
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
			universeMock.EXPECT().
				Look().
				Return(&universe.View{}, nil).
				Times(testCase.lookCalls)
			outMock := mockscreenprinter.NewMockScreenPrinter(ctrl)
			outMock.EXPECT().
				Println(gomock.Any()).
				AnyTimes()
			support := support.Support{
				Out: outMock,
			}
			matcherMock := mockmatch.NewMockMatcher(ctrl)
			matcherMock.EXPECT().Ingest(gomock.Any()).AnyTimes()
			matcherMock.EXPECT().
				Match(gomock.Any()).
				DoAndReturn(func(s string) []string {
					return []string{s}
				}).
				AnyTimes()
			textui := &textUI{
				s:              support,
				u:              universeMock,
				commandMatcher: matcherMock,
			}

			// act
			quit, err := textui.call(testCase.request)

			// assert
			testCase.assert(quit, err)
		})
	}
}

func TestCallGoList(t *testing.T) {
	testCases := []struct {
		name         string
		golistReturn []universe.Neighbor
		golistError  error
		golistCalls  int
		outInput     string
		outCalls     int
		assert       func(bool, error)
	}{
		{
			name:         "go success no params",
			golistReturn: []universe.Neighbor{{Index: 0, Name: "Moon"}},
			golistCalls:  1,
			outInput:     "Moon",
			outCalls:     1,
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
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
				GoList().
				Return(testCase.golistReturn, testCase.golistError).
				Times(testCase.golistCalls)
			outMock := mockscreenprinter.NewMockScreenPrinter(ctrl)
			outMock.EXPECT().
				Println(testCase.outInput).
				Times(testCase.outCalls)
			support := support.Support{
				Out: outMock,
			}
			matcherMock := mockmatch.NewMockMatcher(ctrl)
			matcherMock.EXPECT().Ingest(gomock.Any()).AnyTimes()
			matcherMock.EXPECT().
				Match(gomock.Any()).
				DoAndReturn(func(s string) []string {
					return []string{s}
				}).
				AnyTimes()
			textui := &textUI{
				s:              support,
				u:              universeMock,
				commandMatcher: matcherMock,
			}

			// act
			quit, err := textui.call(&goRequest{})

			// assert
			testCase.assert(quit, err)
		})
	}
}

func TestCallHelp(t *testing.T) {
	testCases := []struct {
		name       string
		request    interface{}
		helpReturn []action.Type
		helpError  error
		helpCalls  int
		outInput   string
		outCalls   int
		assert     func(bool, error)
	}{
		{
			name:    "success 0 params (returning Go)",
			request: &helpRequest{},
			helpReturn: []action.Type{
				action.TypeGo,
			},
			helpCalls: 1,
			outInput:  "Go   - Travel",
			outCalls:  1,
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name:    "success 0 params (returning Look)",
			request: &helpRequest{},
			helpReturn: []action.Type{
				action.TypeLook,
			},
			helpCalls: 1,
			outInput:  "Look - Look around",
			outCalls:  1,
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name:      "help call fail",
			request:   &helpRequest{},
			helpError: errors.New("some go error"),
			helpCalls: 1,
			assert: func(quit bool, err error) {
				assert.Error(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "success 1 param",
			request: &helpRequest{
				Command: "go",
			},
			outInput: "Go <destination> - Travel to destination",
			outCalls: 1,
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "fail 1 unknown param",
			request: &helpRequest{
				Command: "DoAFlip",
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
				Times(testCase.outCalls)
			support := support.Support{
				Out: outMock,
			}
			matcherMock := mockmatch.NewMockMatcher(ctrl)
			matcherMock.EXPECT().Ingest(gomock.Any()).AnyTimes()
			matcherMock.EXPECT().
				Match(gomock.Any()).
				DoAndReturn(func(s string) []string {
					return []string{s}
				}).
				AnyTimes()
			textui := &textUI{
				s:              support,
				u:              universeMock,
				commandMatcher: matcherMock,
			}

			// act
			quit, err := textui.call(testCase.request)

			// assert
			testCase.assert(quit, err)
		})
	}
}

func TestCallLook(t *testing.T) {
	testCases := []struct {
		name         string
		request      interface{}
		lookReturn   *universe.View
		lookError    error
		out1Expected string
		out2Expected string
		outCalls     int
		assert       func(bool, error)
	}{
		{
			name:    "success",
			request: &lookRequest{},
			lookReturn: &universe.View{
				Name:        "Mars",
				Description: "a red planet",
				Path:        []string{"sol", "Mars"},
			},
			out1Expected: "You are at Mars, a red planet.",
			out2Expected: "sol/Mars/",
			outCalls:     1,
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name:      "call failed",
			request:   &lookRequest{},
			lookError: errors.New("some look error"),
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
				Look().
				Return(testCase.lookReturn, testCase.lookError).
				Times(1)
			outMock := mockscreenprinter.NewMockScreenPrinter(ctrl)
			first := outMock.EXPECT().
				Println(testCase.out1Expected).
				Times(testCase.outCalls)
			outMock.EXPECT().
				Println(testCase.out2Expected).
				After(first).
				Times(testCase.outCalls)
			support := support.Support{
				Out: outMock,
			}
			matcherMock := mockmatch.NewMockMatcher(ctrl)
			matcherMock.EXPECT().Ingest(gomock.Any()).AnyTimes()
			matcherMock.EXPECT().
				Match(gomock.Any()).
				DoAndReturn(func(s string) []string {
					return []string{s}
				}).
				AnyTimes()
			textui := &textUI{
				s:              support,
				u:              universeMock,
				commandMatcher: matcherMock,
			}

			// act
			quit, err := textui.call(testCase.request)

			// assert
			testCase.assert(quit, err)
		})
	}
}

func TestCallSell(t *testing.T) {
	testCases := []struct {
		name       string
		request    interface{}
		sellAmount int
		sellItem   string
		sellReturn error
		sellCalls  int
		assert     func(bool, error)
	}{
		{
			name: "sell success",
			request: &sellRequest{
				Amount: 3,
				Item:   "computers",
			},
			sellAmount: 3,
			sellItem:   "computers",
			sellCalls:  1,
			assert: func(quit bool, err error) {
				assert.NoError(t, err)
				assert.Equal(t, false, quit)
			},
		},
		{
			name: "sell call fail",
			request: &sellRequest{
				Amount: 3,
				Item:   "computers",
			},
			sellAmount: 3,
			sellItem:   "computers",
			sellReturn: errors.New("some sell error"),
			sellCalls:  1,
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
				Sell(testCase.sellAmount, testCase.sellItem).
				Return(testCase.sellReturn).
				Times(testCase.sellCalls)
			matcherMock := mockmatch.NewMockMatcher(ctrl)
			matcherMock.EXPECT().Ingest(gomock.Any()).AnyTimes()
			matcherMock.EXPECT().
				Match(gomock.Any()).
				DoAndReturn(func(s string) []string {
					return []string{s}
				}).
				AnyTimes()
			textui := &textUI{
				u:              universeMock,
				commandMatcher: matcherMock,
			}

			// act
			quit, err := textui.call(testCase.request)

			// assert
			testCase.assert(quit, err)
		})
	}
}
