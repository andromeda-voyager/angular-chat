package util

import (
	"crypto/rand"
	"encoding/base64"
	mrand "math/rand"
	"time"
)

func init() {
	mrand.Seed(time.Now().UTC().UnixNano())
}

// GetRandomBytes Returns a random slice of bytes
func GetRandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	for err != nil {
		_, err = rand.Read(b)
	}
	return b
}

// NewRandomSecureString Returns a random cryptographically secure string
func NewRandomSecureString(n int) string {
	b := GetRandomBytes(n)
	return base64.RawStdEncoding.EncodeToString(b)
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// NewRandomString returns a pseudo random alphanumeric string
func NewRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[mrand.Intn(len(chars))]
	}
	return string(b)
}
