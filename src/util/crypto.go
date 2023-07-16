package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(tar string) string {
	b := []byte(tar)
	sha256 := sha256.Sum256(b)
	return hex.EncodeToString(sha256[:])
}
