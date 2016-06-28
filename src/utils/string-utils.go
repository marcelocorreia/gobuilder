package utils

import (
	"log"
	"crypto/rand"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52 possibilities
	letterIdxBits = 6                    // 6 bits to represent 64 possibilities / indexes
	letterIdxMask = 1 << letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

type Strings interface {
	SecureRandomAlphaString(length int) string
	SecureRandomBytes(length int) []byte
}

type StringUtils struct{}

func (su StringUtils)SecureRandomAlphaString(length int) string {
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j % bufferSize == 0 {
			randomBytes = su.SecureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j % length] & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i++
		}
	}

	return string(result)
}

// SecureRandomBytes returns the requested number of bytes using crypto/rand
func (su StringUtils) SecureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}