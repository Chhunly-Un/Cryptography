package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

// SHA1Hex returns the SHA1 hex string of the given input.
func SHA1Hex(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
