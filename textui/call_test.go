package textui

import (
	"corsairtext/e"
	"corsairtext/support"
	"corsairtext/support/screenprinter/mockscreenprinter"
	"corsairtext/textui/match/mockmatch"
	"corsairtext/universe"
	"corsairtext/universe/mockuniverse"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	testCases := []struct {
		name    string
		command string
		assert  func(error)
	}{
		{
			name:    "success quit",
			command: "qui",
			assert: func(err error) {
				assert.Error(t, err)
				assert.True(t, e.IsQuitError(err))
			},
		},
		{
			name:    "fail bad command",
			command: "abcxyz",
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
		{
			name:    "fail empty command",
			command: "",
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			textui := &textUI{
				commandMatcher: NewCommandMatcher(),
			}

			// act
			err := textui.act(testCase.command)

			// assert
			testCase.assert(err)
		})
	}
}

func TestCallBuy(t *testing.T) {
	testCases := []struct {
		name      string
		command   string
		buyAmount int
		buyItem   string
		buyReturn error
		buyCalls  int
		assert    func(error)
	}{
		{
			name:      "buy success",
			command:   "BUY 3 computers",
			buyAmount: 3,
			buyItem:   "computers",
			buyCalls:  1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name:      "buy call fail",
			command:   "BUY 3 computers",
			buyAmount: 3,
			buyItem:   "computers",
			buyReturn: errors.New("some buy error"),
			buyCalls:  1,
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
		{
			name:    "buy missing param",
			command: "BUY 3",
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
		{
			name:    "buy bad param",
			command: "BUY three computers",
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
			textui := &textUI{
				a:              actionMock,
				commandMatcher: NewCommandMatcher(),
			}

			// act
			err := textui.act(testCase.command)

			// assert
			testCase.assert(err)
		})
	}
}

func TestCallGo(t *testing.T) {
	testCases := []struct {
		name               string
		command            string
		goDestination      string
		goReturn           error
		goCalls            int
		localLocationCalls int
		assert             func(error)
	}{
		{
			name:               "go success with dest",
			command:            "gO mars",
			goDestination:      "mars",
			goCalls:            1,
			localLocationCalls: 1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name:          "go call fail",
			command:       "go mars",
			goDestination: "mars",
			goReturn:      errors.New("some go error"),
			goCalls:       1,
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
		{
			name:    "go fail too many params",
			command: "go mars jupiter",
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
				commandMatcher:  NewCommandMatcher(),
				locationMatcher: matcherMock,
			}

			// act
			err := textui.act(testCase.command)

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
				commandMatcher: NewCommandMatcher(),
			}

			// act
			err := textui.act("G")

			// assert
			testCase.assert(err)
		})
	}
}

func TestCallHelp(t *testing.T) {
	testCases := []struct {
		name            string
		command         string
		listLocalReturn map[string]interface{}
		listLocalCalls  int
		outInput        string
		outCalls        int
		assert          func(error)
	}{
		// {
		// 	name:    "success 0 params (returning dig)",
		// 	command: "help",
		// 	listLocalReturn: map[string]interface{}{
		// 		CommandDig: nil,
		// 	},
		// 	listLocalCalls: 1,
		// 	outInput:       "dig - Mine for ore",
		// 	outCalls:       1,
		// 	assert: func(err error) {
		// 		assert.NoError(t, err)
		// 	},
		// },
		// {
		// 	name:    "success 0 params (returning Look)",
		// 	command: "hel",
		// 	listLocalReturn: map[string]interface{}{
		// 		CommandLook: nil,
		// 	},
		// 	listLocalCalls: 1,
		// 	outInput:       "look - Look around",
		// 	outCalls:       1,
		// 	assert: func(err error) {
		// 		assert.NoError(t, err)
		// 	},
		// },
		{
			name:     "success 1 param",
			command:  "help go",
			outInput: "go <destination> - Travel to destination",
			outCalls: 1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name:    "fail 1 unknown param",
			command: "help DoAFlip",
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
				ListLocalActions().
				Return(testCase.listLocalReturn).
				Times(testCase.listLocalCalls)
			outMock := mockscreenprinter.NewMockScreenPrinter(ctrl)
			outMock.EXPECT().
				Println(testCase.outInput).
				Times(testCase.outCalls)
			support := support.Support{
				Out: outMock,
			}
			textui := &textUI{
				s:              support,
				i:              informationMock,
				commandMatcher: NewCommandMatcher(),
			}
			commandDescriptionMap[CommandHelp].Handler = helpHandlerTableEntry

			// act
			err := textui.act(testCase.command)

			// assert
			testCase.assert(err)
		})
	}
}

func TestCallInventory(t *testing.T) {
	testCases := []struct {
		name            string
		command         string
		inventoryReturn universe.Ship
		inventoryCalls  int
		out1Expected    string
		out2Expected    string
		out3Expected    string
		out4Expected    string
		outCalls        int
		assert          func(error)
	}{
		{
			name:    "success",
			command: "inventory",
			inventoryReturn: universe.Ship{
				Money:        12,
				ItemCapacity: 3,
				Items: map[string]*universe.ItemLot{
					"ore": {
						Count:    3,
						UnitCost: 1,
						Origin:   "Tranquility",
					},
				},
			},
			inventoryCalls: 1,
			out1Expected:   "Money: 12",
			out2Expected:   "Capacity: 3",
			out3Expected:   "Load: 3",
			out4Expected:   " ore: 3",
			outCalls:       1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name:    "too many parameters",
			command: "inventory bluck",
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
				Inventory().
				Return(testCase.inventoryReturn).
				Times(testCase.inventoryCalls)
			outMock := mockscreenprinter.NewMockScreenPrinter(ctrl)
			first := outMock.EXPECT().
				Println(testCase.out1Expected).
				Times(testCase.outCalls)
			second := outMock.EXPECT().
				Println(testCase.out2Expected).
				After(first).
				Times(testCase.outCalls)
			third := outMock.EXPECT().
				Println(testCase.out3Expected).
				After(second).
				Times(testCase.outCalls)
			outMock.EXPECT().
				Println(testCase.out4Expected).
				After(third).
				Times(testCase.outCalls)
			support := support.Support{
				Out: outMock,
			}
			textui := &textUI{
				s:              support,
				i:              informationMock,
				commandMatcher: NewCommandMatcher(),
			}

			// act
			err := textui.act(testCase.command)

			// assert
			testCase.assert(err)
		})
	}
}

func TestCallLook(t *testing.T) {
	testCases := []struct {
		name                string
		command             string
		localLocationReturn *universe.View
		localLocationCalls  int
		out1Expected        string
		out2Expected        string
		outCalls            int
		assert              func(error)
	}{
		{
			name:    "success",
			command: "lOOk",
			localLocationReturn: &universe.View{
				Name:        "Mars",
				Description: "a red planet",
				Path:        []string{"sol", "Mars"},
			},
			localLocationCalls: 1,
			out1Expected:       "You are at Mars, a red planet.",
			out2Expected:       "sol/Mars/",
			outCalls:           1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name:    "extra parameters",
			command: "look around",
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
				LocalLocation().
				Return(testCase.localLocationReturn).
				Times(testCase.localLocationCalls)
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
			textui := &textUI{
				s:              support,
				i:              informationMock,
				commandMatcher: NewCommandMatcher(),
			}

			// act
			err := textui.act(testCase.command)

			// assert
			testCase.assert(err)
		})
	}
}

func TestCallSell(t *testing.T) {
	testCases := []struct {
		name       string
		command    string
		sellAmount int
		sellItem   string
		sellReturn error
		sellCalls  int
		assert     func(error)
	}{
		{
			name:       "sell success",
			command:    "sell 3 computers",
			sellAmount: 3,
			sellItem:   "computers",
			sellCalls:  1,
			assert: func(err error) {
				assert.NoError(t, err)
			},
		},
		{
			name:       "sell call fail",
			command:    "sell 3 computers",
			sellAmount: 3,
			sellItem:   "computers",
			sellReturn: errors.New("some sell error"),
			sellCalls:  1,
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
		{
			name:    "sell missing param",
			command: "SELL 3",
			assert: func(err error) {
				assert.Error(t, err)
			},
		},
		{
			name:    "sell bad param",
			command: "sell three computers",
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
				commandMatcher: NewCommandMatcher(),
			}

			// act
			err := textui.act(testCase.command)

			// assert
			testCase.assert(err)
		})
	}
}
