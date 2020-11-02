package textui

import (
	"corsairtext/support"
	"corsairtext/universe"
)

// TextUI is the entry interface for the text ui
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type TextUI interface {
	// Run starts the text based UI
	Run()
}

// NewTextUI create a new text ui
func NewTextUI(s support.Support, u universe.Universe) TextUI {
	return &textUI{
		s: s,
		u: u,
	}
}

// textUI is the concrete implementation of TextUI
type textUI struct {
	s support.Support
	u universe.Universe
}

// Run is the main text ui entry point
func (t *textUI) Run() {
	t.act("look", t.u.Actions())
	for {
		t.s.Out.Print(" ready> ")
		text, err := t.s.In.Readln()
		if err != nil {
			return
		}
		quit, err := t.act(text, t.u.Actions())
		if quit {
			break
		}
		if err != nil {
			t.s.Out.Println("Error: ", err)
		}
	}
}
