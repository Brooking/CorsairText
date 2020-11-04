package textui

import (
	"corsairtext/action"
	"corsairtext/e"
	"corsairtext/support"
	"corsairtext/universe"

	"strings"

	"github.com/pkg/errors"
)

// TextUI is the entry interface for the text ui
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type TextUI interface {
	// Run starts the text based UI
	Run()
}

// NewTextUI create a new text ui
func NewTextUI(s support.Support, u universe.Action) TextUI {
	return &textUI{
		s: s,
		u: u,
	}
}

// textUI is the concrete implementation of TextUI
type textUI struct {
	s support.Support
	u universe.Action
}

// Run is the main text ui entry point
func (t *textUI) Run() {
	t.act("look")
	for {
		t.s.Out.Print("ready> ")
		text, err := t.s.In.Readln()
		t.s.Out.Println("")
		if err != nil {
			return
		}
		quit, err := t.act(text)
		if quit {
			break
		}
		if err != nil {
			cause := errors.Cause(err)

			switch {
			case e.IsShowToUserError(cause):
				t.s.Out.Println(cause.Error())
			default:
				t.s.Out.Println(strings.Join([]string{"Error: ", err.Error()}, ""))
			}

			switch {
			case e.IsShowAllHelpError(cause):
				t.call(Request{Type: action.TypeHelp})
			case e.IsShowHelpError(cause):
				t.call(Request{
					Type:       action.TypeHelp,
					Parameters: []interface{}{e.GetActionTypeForHelp(cause).String()},
				})
			case e.IsShowAdjacencyError(cause):
				t.call(Request{Type: action.TypeGoList})
			}
		}
	}
}

// act handles a user's command
func (t *textUI) act(command string) (bool, error) {
	request, err := t.parseAction(command)
	if err != nil {
		return false, errors.Wrap(err, "parsing action")
	}
	return t.call(request)
}
