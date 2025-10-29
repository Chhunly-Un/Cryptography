package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

// SHA512Hex returns the SHA-512 hex string of the given input (lowercase).
func SHA512Hex(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
