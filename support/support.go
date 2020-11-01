package support

import (
	"corsairtext/support/keyboardreader"
	"corsairtext/support/random"
	"corsairtext/support/screenprinter"
)

// Support contains all of the mockable interfaces that we get from outside
type Support struct {
	In   keyboardreader.KeyboardReader
	Out  screenprinter.ScreenPrinter
	Rand random.Random
}

// NewSupport creates a real support struct
func NewSupport() Support {
	return Support{
		In:   keyboardreader.NewKeyboardReader(),
		Out:  screenprinter.NewScreenPrinter(),
		Rand: random.NewRandom(),
	}
}
