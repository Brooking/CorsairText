package textui

import (
	"corsairtext/support"
	"corsairtext/universe"
)

// TextUI is the entry interface for the text ui
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
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
		spot := t.u.WhereAmI()
		t.s.Out.Println("You are at", spot.Description())
		t.s.Out.Println(spot.Path())
		t.s.Out.Print("ready> ")
		text, err := t.s.In.ReadLn()
		if err != nil {
			return
		}
		if text == "q" {
			return
		}
		t.s.Out.Println("'" + text + "'")
	}
}
