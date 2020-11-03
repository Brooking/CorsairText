package screenprinter

import (
	"fmt"

	"github.com/pkg/errors"
)

// ScreenPrinter is the interface abstracting printing to the screen
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type ScreenPrinter interface {
	// Print places text on the screen without a terminating \n
	Print(a ...interface{})

	// Println places text on the screen with a terminating \n
	Println(a string)
}

// NewScreenPrinter creates a new screen printer
func NewScreenPrinter() ScreenPrinter {
	return screenPrinter{}
}

// screenPrinter implements a ScreenPrinter interface to the real screen
type screenPrinter struct{}

// Print places text on the screen without a terminating \n
func (p screenPrinter) Print(a ...interface{}) {
	_, err := fmt.Print(a...)
	if err != nil {
		panic(errors.Wrap(err, "fmt.Print"))
	}
}

// Println places text on the screen with a terminating \n
func (p screenPrinter) Println(a string) {
	_, err := fmt.Println(a)
	if err != nil {
		panic(errors.Wrap(err, "fmt.Println"))
	}
}
