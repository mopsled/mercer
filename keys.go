package mercer

import (
	"crypto/rand"
)

func PrivateKey() (key []byte, err error) {
	key = make([]byte, 32)
	_, err = rand.Read(key)
	return
}
