package main

import (
	"corsairtext/support"
	"corsairtext/textui"
	"corsairtext/universe"
)

// main is the entry point for corsair text
func main() {
	s := support.NewSupport()
	u := universe.NewUniverse(s)
	ui := textui.NewTextUI(s, u)
	ui.Run()
}
