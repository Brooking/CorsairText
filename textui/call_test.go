package textui

import (
	"errors"
	"testing"

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
		command interface{}
		assert  func(error)
	}{
		{
			name:    "success quit",
			command: &quitCommand{},
			assert: func(err error) {
				assert.Error(t, err)
				assert.True(t, e.IsQuitError(err))
			},
		},
		{
			name:    "fail bad struct",
			command: e.AmbiguousCommandError{},
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
		{
			name:    "fail nil struct",
			command: nil,
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			textui := &textUI{}

			// act
			err := textui.call(testCase.command)

			// assert
			testCase.assert(err)
		})
	}
}

func TestCallBuy(t *testing.T) {
	testCases := []struct {
		name      string
		command   interface{}
		buyAmount int
		buyItem   string
		buyReturn error
		buyCalls  int
		assert    func(error)
	}{
		{
			name: "buy success",
			command: &buyCommand{
				Amount: 3,
				Item:   "computers",
			},
			buyAmount: 3,
			buyItem:   "computers",
			buyCalls:  1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "buy call fail",
			command: &buyCommand{
				Amount: 3,
				Item:   "computers",
			},
			buyAmount: 3,
			buyItem:   "computers",
			buyReturn: errors.New("some buy error"),
			buyCalls:  1,
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			actionMock := mockuniverse.NewMockAction(ctrl)
			actionMock.EXPECT().
				Buy(testCase.buyAmount, testCase.buyItem).
				Return(testCase.buyReturn).
				Times(testCase.buyCalls)
			matcherMock := mockmatch.NewMockMatcher(ctrl)
			matcherMock.EXPECT().
				Match(gomock.Any()).
				DoAndReturn(func(s string) []string {
					return []string{s}
				}).
				AnyTimes()
			textui := &textUI{
				a:              actionMock,
				commandMatcher: matcherMock,
			}

			// act
			err := textui.call(testCase.command)

			// assert
			testCase.assert(err)
		})
	}
}

func TestCallGo(t *testing.T) {
	testCases := []struct {
		name               string
		command            interface{}
		goDestination      string
		goReturn           error
		goCalls            int
		localLocationCalls int
		assert             func(error)
	}{
		{
			name: "go success with dest",
			command: &goCommand{
				Destination: "mars",
			},
			goDestination:      "mars",
			goCalls:            1,
			localLocationCalls: 1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "go call fail",
			command: &goCommand{
				Destination: "mars",
			},
			goDestination: "mars",
			goReturn:      errors.New("some go error"),
			goCalls:       1,
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			actionMock := mockuniverse.NewMockAction(ctrl)
			actionMock.EXPECT().
				Go(testCase.goDestination).
				Return(testCase.goReturn).
				Times(testCase.goCalls)
			informationMock := mockuniverse.NewMockInformation(ctrl)
			informationMock.EXPECT().
				LocalLocation().
				Return(&universe.View{}).
				Times(testCase.localLocationCalls)
			informationMock.EXPECT().
				ListAdjacentLocations().
				Return([]string{"mars"}).
				AnyTimes()
			outMock := mockscreenprinter.NewMockScreenPrinter(ctrl)
			outMock.EXPECT().
				Println(gomock.Any()).
				AnyTimes()
			support := support.Support{
				Out: outMock,
			}
			matcherMock := mockmatch.NewMockMatcher(ctrl)
			matcherMock.EXPECT().
				Match(gomock.Any()).
				DoAndReturn(func(s string) []string {
					return []string{s}
				}).
				AnyTimes()
			textui := &textUI{
				s:               support,
				a:               actionMock,
				i:               informationMock,
				commandMatcher:  matcherMock,
				locationMatcher: matcherMock,
			}

			// act
			err := textui.call(testCase.command)

			// assert
			testCase.assert(err)
		})
	}
}

func TestCallGoList(t *testing.T) {
	testCases := []struct {
		name         string
		golistReturn []string
		golistCalls  int
		outInput     string
		outCalls     int
		assert       func(error)
	}{
		{
			name:         "go success no params",
			golistReturn: []string{"Moon"},
			golistCalls:  1,
			outInput:     "Moon",
			outCalls:     1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			informationMock := mockuniverse.NewMockInformation(ctrl)
			informationMock.EXPECT().
				ListAdjacentLocations().
				Return(testCase.golistReturn).
				Times(testCase.golistCalls)
			outMock := mockscreenprinter.NewMockScreenPrinter(ctrl)
			outMock.EXPECT().
				Println(testCase.outInput).
				Times(testCase.outCalls)
			support := support.Support{
				Out: outMock,
			}
			matcherMock := mockmatch.NewMockMatcher(ctrl)
			matcherMock.EXPECT().
				Match(gomock.Any()).
				DoAndReturn(func(s string) []string {
					return []string{s}
				}).
				AnyTimes()
			textui := &textUI{
				s:              support,
				i:              informationMock,
				commandMatcher: matcherMock,
			}

			// act
			err := textui.call(&goCommand{})

			// assert
			testCase.assert(err)
		})
	}
}

func TestCallHelp(t *testing.T) {
	testCases := []struct {
		name            string
		command         interface{}
		listLocalReturn map[string]interface{}
		listLocalCalls  int
		outInput        string
		outCalls        int
		assert          func(error)
	}{
		{
			name:    "success 0 params (returning go)",
			command: &helpCommand{},
			listLocalReturn: map[string]interface{}{
				CommandGo: nil,
			},
			listLocalCalls: 1,
			outInput:       "go   - Travel",
			outCalls:       1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name:    "success 0 params (returning Look)",
			command: &helpCommand{},
			listLocalReturn: map[string]interface{}{
				CommandLook: nil,
			},
			listLocalCalls: 1,
			outInput:       "look - Look around",
			outCalls:       1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "success 1 param",
			command: &helpCommand{
				Command: CommandGo,
			},
			outInput: "go <destination> - Travel to destination",
			outCalls: 1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "fail 1 unknown param",
			command: &helpCommand{
				Command: "DoAFlip",
			},
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			informationMock := mockuniverse.NewMockInformation(ctrl)
			informationMock.EXPECT().
				ListLocalCommands().
				Return(testCase.listLocalReturn).
				Times(testCase.listLocalCalls)
			outMock := mockscreenprinter.NewMockScreenPrinter(ctrl)
			outMock.EXPECT().
				Println(testCase.outInput).
				Times(testCase.outCalls)
			support := support.Support{
				Out: outMock,
			}
			matcherMock := mockmatch.NewMockMatcher(ctrl)
			matcherMock.EXPECT().
				Match(gomock.Any()).
				DoAndReturn(func(s string) []string {
					return []string{s}
				}).
				AnyTimes()
			textui := &textUI{
				s:              support,
				i:              informationMock,
				commandMatcher: matcherMock,
			}

			// act
			err := textui.call(testCase.command)

			// assert
			testCase.assert(err)
		})
	}
}

func TestCallLook(t *testing.T) {
	testCases := []struct {
		name                string
		command             interface{}
		localLocationReturn *universe.View
		out1Expected        string
		out2Expected        string
		outCalls            int
		assert              func(error)
	}{
		{
			name:    "success",
			command: &lookCommand{},
			localLocationReturn: &universe.View{
				Name:        "Mars",
				Description: "a red planet",
				Path:        []string{"sol", "Mars"},
			},
			out1Expected: "You are at Mars, a red planet.",
			out2Expected: "sol/Mars/",
			outCalls:     1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			informationMock := mockuniverse.NewMockInformation(ctrl)
			informationMock.EXPECT().
				LocalLocation().
				Return(testCase.localLocationReturn).
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
			matcherMock.EXPECT().
				Match(gomock.Any()).
				DoAndReturn(func(s string) []string {
					return []string{s}
				}).
				AnyTimes()
			textui := &textUI{
				s:              support,
				i:              informationMock,
				commandMatcher: matcherMock,
			}

			// act
			err := textui.call(testCase.command)

			// assert
			testCase.assert(err)
		})
	}
}

func TestCallSell(t *testing.T) {
	testCases := []struct {
		name       string
		command    interface{}
		sellAmount int
		sellItem   string
		sellReturn error
		sellCalls  int
		assert     func(error)
	}{
		{
			name: "sell success",
			command: &sellCommand{
				Amount: 3,
				Item:   "computers",
			},
			sellAmount: 3,
			sellItem:   "computers",
			sellCalls:  1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "sell call fail",
			command: &sellCommand{
				Amount: 3,
				Item:   "computers",
			},
			sellAmount: 3,
			sellItem:   "computers",
			sellReturn: errors.New("some sell error"),
			sellCalls:  1,
			assert: func(err error) {
				assert.Error(t, err)
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
			matcherMock.EXPECT().
				Match(gomock.Any()).
				DoAndReturn(func(s string) []string {
					return []string{s}
				}).
				AnyTimes()
			textui := &textUI{
				a:              universeMock,
				commandMatcher: matcherMock,
			}

			// act
			err := textui.call(testCase.command)

			// assert
			testCase.assert(err)
		})
	}
}
