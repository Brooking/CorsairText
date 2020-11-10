package textui

import (
	"corsairtext/e"
	"corsairtext/match"
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
func NewTextUI(s support.Support, a universe.Action, i universe.Information) TextUI {
	return &textUI{
		s:               s,
		a:               a,
		i:               i,
		commandMatcher:  NewCommandMatcher(),
		locationMatcher: match.NewMatcher(i.ListLocations(), false),
	}
}

// textUI is the concrete implementation of TextUI
type textUI struct {
	s               support.Support
	a               universe.Action
	i               universe.Information
	commandMatcher  match.Matcher
	locationMatcher match.Matcher
}

// Run is the main text ui entry point
func (t *textUI) Run() {
	var err error
	err = t.call(&lookCommand{})
	if err != nil {
		t.showError(err)
	}
	for {
		t.s.Out.Print("ready> ")
		text, err := t.s.In.Readln()
		t.s.Out.Println("")
		if err != nil {
			return
		}

		err = t.act(text)
		if e.IsQuitError(err) {
			break
		}

		if err != nil {
			t.showError(err)
		}
	}
}

// act handles a user's command
func (t *textUI) act(commandString string) error {
	command, err := t.parse(commandString)
	if err != nil {
		return errors.Wrap(err, "parsing action")
	}
	return t.call(command)
}

func (t *textUI) showError(err error) {
	cause := errors.Cause(err)
	switch {
	case e.IsShowToUserError(cause):
		t.s.Out.Println(cause.Error())
	default:
		t.s.Out.Println(strings.Join([]string{"Error: ", err.Error()}, ""))
	}

	switch {
	case e.IsShowAllHelpError(cause):
		t.call(helpCommand{})
	case e.IsShowHelpError(cause):
		t.call(helpCommand{
			Command: e.GetCommandForHelp(cause),
		})
	case e.IsShowAdjacencyError(cause):
		t.call(goCommand{})
	}
}
