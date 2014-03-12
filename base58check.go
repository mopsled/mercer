package mercer

import (
	"crypto/sha256"
)

func Base58CheckEncode(bytes []byte) string {
	bytes = append([]byte{0x80}, bytes...)
	checksum := ChecksumBytes(bytes)
	bytes = append(bytes, checksum...)
	return Encode(bytes)
}

func ChecksumBytes(bytes []byte) []byte {
	shaFirst := sha256.Sum256(bytes)
	doubleSha256 := sha256.Sum256(shaFirst[:])
	return doubleSha256[:4]
}
