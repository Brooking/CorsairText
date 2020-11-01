package screenprinter

import "fmt"

// ScreenPrinter is the interface abstracting printing to the screen
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type ScreenPrinter interface {
	Print(a ...interface{}) (n int, err error)
	Println(a ...interface{}) (n int, err error)
}

// NewScreenPrinter creates a new screen printer
func NewScreenPrinter() ScreenPrinter {
	return screenPrinter{}
}

// screenPrinter implements a ScreenPrinter interface to the real screen
type screenPrinter struct{}

// Print places text on the screen without a terminating \n
func (p screenPrinter) Print(a ...interface{}) (n int, err error) {
	return fmt.Print(a...)
}

// Println places text on the screen with a terminating \n
func (p screenPrinter) Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}
