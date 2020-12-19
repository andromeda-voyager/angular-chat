package util

import (
	"crypto/rand"
	"encoding/base64"
)

// GetRandomBytes Returns a random slice of bytes
func GetRandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	for err != nil {
		_, err = rand.Read(b)
	}
	return b
}

// NewRandomString Returns a random string
func NewRandomString(n int) string {
	b := GetRandomBytes(n)
	return base64.RawStdEncoding.EncodeToString(b)
}
