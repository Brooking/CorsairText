package main

import (
	"corsairtext/support"
	"corsairtext/support/keyboardreader"
	"corsairtext/support/screenprinter"
	"corsairtext/textui"
	"corsairtext/universe"
)

// main is the entry point for corsair text
func main() {
	s := support.Support{
		In:  keyboardreader.NewKeyboardReader(),
		Out: screenprinter.NewScreenPrinter(),
	}
	u := universe.NewUniverse(s)
	ui := textui.NewTextUI(s, u)
	ui.Run()
}
