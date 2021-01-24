package random

import (
	"crypto/rand"
	"encoding/base64"
	mrand "math/rand"
	"time"
)

func init() {
	mrand.Seed(time.Now().UTC().UnixNano())
}

// GetSecureBytes Returns a random slice of bytes
func GetSecureBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	for err != nil {
		_, err = rand.Read(b)
	}
	return b
}

// NewSecureString Returns a random cryptographically secure string
func NewSecureString(n int) string {
	b := GetSecureBytes(n)
	return base64.RawStdEncoding.EncodeToString(b)
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// // NewString returns a pseudo random alphanumeric string
// func NewString(n int) string {
// 	b := make([]byte, n)
// 	for i := range b {
// 		b[i] = chars[mrand.Intn(len(chars))]
// 	}
// 	return string(b)
// }
