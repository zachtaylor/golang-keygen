package keygen

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// NewCrypto uses crypto/rand
func NewCrypto(len int) string {
	b := make([]byte, len)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
