package mercer

import (
	"bytes"
	"math/big"
)

func Encode(b []byte) string {
	base58characters := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	var buffer bytes.Buffer

	x := new(big.Int).SetBytes(b)
	for x.Cmp(big.NewInt(0)) == 1 {
		var remainder *big.Int
		x, remainder = x.DivMod(x, big.NewInt(58), new(big.Int))
		buffer.WriteByte(base58characters[remainder.Uint64()])
	}

	leadingZeros := leadingZerosForBytes(b)
	buffer.WriteString(leadingZeros)
	return reverse(buffer.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func leadingZerosForBytes(b []byte) string {
	var buffer bytes.Buffer
	for _, v := range b {
		if v == 0 {
			buffer.WriteString("1")
		} else {
			break
		}
	}
	return buffer.String()
}
