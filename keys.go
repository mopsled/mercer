package mercer

import (
	"crypto/rand"
	"crypto/sha256"
)

func PrivateKey() (key []byte, err error) {
	key = make([]byte, 64)
	_, err = rand.Read(key)
	return
}

func ChecksumBytes(bytes []byte) []byte {
	shaFirst := sha256.Sum256(bytes)
	doubleSha256 := sha256.Sum256(shaFirst[:])
	return doubleSha256[:4]
}
