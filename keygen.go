package keygen

// Func is a func() string
type Func = func() string

// Default uses New(32)
func Default() string { return New(32) }

// NewFunc returns a Func that calls New with len
func NewFunc(len int) Func {
	return func() string {
		return New(len)
	}
}

// NewFunc returns a Func that calls NewCrypto with len
func NewCryptoFunc(len int) Func {
	return func() string {
		return NewCrypto(len)
	}
}
