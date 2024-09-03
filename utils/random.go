package util

import (
	"crypto/rand"
	"encoding/hex"
)

func generatePass(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return hex.EncodeToString(b)
}
