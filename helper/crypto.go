package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256toString(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	result := hex.EncodeToString(h.Sum(nil))

	return result
}
