package random

import (
	"crypto/rand"
	"encoding/base64"
	mrand "math/rand"
	"strings"
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

func NewRandomCode(codeLength int) string {
	var b strings.Builder
	for i := 0; i < codeLength; i++ {
		randomIndex := mrand.Intn(len(chars))
		b.WriteByte(chars[randomIndex])
	}
	return b.String()
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
