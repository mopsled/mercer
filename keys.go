package mercer

import (
	"crypto/rand"
)

func PrivateKey() (key []byte, err error) {
	key = make([]byte, 64)
	_, err = rand.Read(key)
	return
}
