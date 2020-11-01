package textui

import (
	"corsairtext/universe"
	"fmt"
)

// TextUI is the entry interface for the text ui
type TextUI interface {
	Run()
}

// NewTextUI create a new text ui
func NewTextUI(u universe.Universe) TextUI {
	return &textUI{
		u: u,
	}
}

// textUI is the concrete implementation of TextUI
type textUI struct {
	u universe.Universe
}

// Run is the main text ui entry point
func (t *textUI) Run() {
	for {
		fmt.Print("ready>")
	}
}
