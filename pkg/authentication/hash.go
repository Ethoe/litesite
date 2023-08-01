package authentication

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
