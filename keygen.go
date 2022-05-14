package keygen

import (
	"math/rand"
	"time"
)

var newRand = rand.NewSource(time.Now().UnixNano()) // allows New to mirror NewCrypto

// NewCryptoFunc creates a Func that wraps NewCrypto(len)
func NewCryptoFunc(len int) Func {
	return func() string {
		return NewCrypto(len)
	}
}

// Func is a func() string
type Func = func() string

// NewFunc returns a Func that wraps New(len)
func NewFunc(len int) Func {
	return func() string {
		return New(len)
	}
}

const defaultNewLen = 21 // evenly distribute 2*rand.int63() [126/6 = 21]

// Default wraps New
func Default() string { return New(defaultNewLen) }

// DefaultFunc creates a Func that wraps Default
func DefaultFunc() func() string {
	return func() string {
		return Default()
	}
}
