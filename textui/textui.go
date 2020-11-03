package textui

import (
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
		t.s.Out.Print(" ready> ")
		text, err := t.s.In.Readln()
		if err != nil {
			return
		}
		quit, err := t.act(text)
		if quit {
			break
		}
		if err != nil {
			t.s.Out.Println(strings.Join([]string{"Error: ", err.Error()}, ""))
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
