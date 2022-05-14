package keygen

import "unsafe"

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

// NotBase64 is a base64-like encoding table, except '+' and '/' are '-' and '_'
const NotBase64 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// New returns a random string of the given length
func New(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, newRand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = newRand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(NotBase64) {
			b[i] = NotBase64[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
