package utils

import (
	"crypto/sha1" //nolint:gosec
	"encoding/hex"
)

func GetPasswordHash(password string) string {
	h := sha1.New() //nolint:gosec
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}
