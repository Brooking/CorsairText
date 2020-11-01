package textui

import (
	"corsairtext/support"
	"corsairtext/universe"
)

// TextUI is the entry interface for the text ui
type TextUI interface {
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
	for {
		t.s.Out.Print("ready>")
		text, err := t.s.In.Read()
		if err != nil {
			return
		}
		t.s.Out.Println(text)
	}
}