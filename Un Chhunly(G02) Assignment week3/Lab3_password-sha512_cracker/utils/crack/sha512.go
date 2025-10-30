package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

func SHA512Hex(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
