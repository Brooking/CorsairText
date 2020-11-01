package main

import (
	"corsairtext/textui"
	"corsairtext/universe"
)

// main is the entry point for corsair text
func main() {
	u := universe.NewUniverse()
	ui := textui.NewTextUI(u)
	ui.Run()
}
