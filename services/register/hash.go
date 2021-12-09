package register

import (
	"encoding/hex"

	sha3 "golang.org/x/crypto/sha3"
)

func Hash(login, password string) string {
	l := []byte(login)
	p := []byte(password)
	lh := make([]byte, 64)
	ph := make([]byte, 64)
	sha3.ShakeSum256(lh, l)
	sha3.ShakeSum256(ph, p)
	hash := append(lh, ph...)
	h := make([]byte, 64)
	sha3.ShakeSum256(h, hash)
	return hex.EncodeToString(h)
}
