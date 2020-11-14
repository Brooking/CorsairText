package support

import (
	"corsairtext/support/keyboardreader"
	"corsairtext/support/random"
	"corsairtext/support/screenprinter"
)

// SupportStruct contains all of the mockable interfaces that we get from outside
type SupportStruct struct {
	In   keyboardreader.KeyboardReader
	Out  screenprinter.ScreenPrinter
	Rand random.Random
}

// NewSupportStruct creates a real support struct
func NewSupportStruct() *SupportStruct {
	return &SupportStruct{
		In:   keyboardreader.NewKeyboardReader(),
		Out:  screenprinter.NewScreenPrinter(),
		Rand: random.NewRandom(),
	}
}
