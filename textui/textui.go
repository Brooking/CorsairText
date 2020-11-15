package textui

import (
	"corsairtext/e"
	"corsairtext/support"
	"corsairtext/textui/commandprocessor"
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
func NewTextUI(s *support.SupportStruct, a universe.Action, i universe.Information) TextUI {
	return &textUI{
		s:  s,
		cp: commandprocessor.NewCommandProcessor(s, a, i),
	}
}

// textUI is the concrete implementation of TextUI
type textUI struct {
	s  *support.SupportStruct
	cp commandprocessor.CommandProcessor
}

// Run is the main text ui entry point
func (t *textUI) Run() {
	var err error
	err = t.cp.ShowLook()
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

		err = t.cp.Obey(text)
		if e.IsQuitError(err) {
			break
		}

		if err != nil {
			t.showError(err)
		}
	}
}

// showError handles displaying errors and any extra needed data
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
		t.cp.ShowAllHelp()
	case e.IsShowHelpError(cause):
		t.cp.ShowHelp(e.GetCommandForHelp(cause))
	case e.IsShowAdjacencyError(cause):
		t.cp.ShowAdjacency()
	}
}
