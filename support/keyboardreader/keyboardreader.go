package keyboardreader

import (
	"bufio"
	"os"
)

// KeyboardReader is an interface that abstracts reading from the keyboard
type KeyboardReader interface {
	Read() (string, error)
}

// NewKeyboardReader creates a new keyboard reader
func NewKeyboardReader() KeyboardReader {
	return read{
		stdin: bufio.NewReader(os.Stdin),
	}
}

// read implements a keyboard reader
type read struct {
	stdin *bufio.Reader
}

// Read returns a single line from the keyboard
func (r read) Read() (string, error) {
	return r.stdin.ReadString('\n')
}
