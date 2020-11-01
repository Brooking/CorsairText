package support

import (
	"corsairtext/support/keyboardreader"
	"corsairtext/support/screenprinter"
)

// Support contains all of the mockable interfaces that we get from outside
type Support struct {
	In  keyboardreader.KeyboardReader
	Out screenprinter.ScreenPrinter
}
