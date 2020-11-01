package random

import "math/rand"

// Random is the interface abstracting random number generation
//go:generate ${GOPATH}/bin/mockgen -destination ./mock${GOPACKAGE}/${GOFILE} -package=mock${GOPACKAGE} -source=${GOFILE}
type Random interface {
	// Intn returns a non-negative random number in [0,n)
	Intn(n int) int
}

// NewRandom creates a new Random interface
func NewRandom() Random {
	return &random{}
}

// random implements the real random generator
type random struct{}

// Intn returns a non-negative random number in [0,n)
func (r *random) Intn(n int) int {
	return rand.Intn(n)
}
