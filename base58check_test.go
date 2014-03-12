package mercer

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestChecksumBytesReturnsExpectedChecksum(t *testing.T) {
	key, err := hex.DecodeString("800C28FCA386C7A227600B2FE50B7CAE11EC86D3BF1FBE471BE89827E19D72AA1D")
	if err != nil {
		t.Fatalf("Error creating key from hex string")
	}
	checksum := ChecksumBytes(key)
	expectedChecksum, err := hex.DecodeString("507A5B8D")
	if err != nil {
		t.Fatalf("Error creating expected checksum from hex string")
	}
	if !bytes.Equal(checksum, expectedChecksum) {
		t.Errorf("Checksum for private key should be %s, but it's %s", hex.EncodeToString(expectedChecksum),
			hex.EncodeToString(checksum))
	}
}

func TestBase58CheckReturnsExpectedString(t *testing.T) {
	key, err := hex.DecodeString("0C28FCA386C7A227600B2FE50B7CAE11EC86D3BF1FBE471BE89827E19D72AA1D")
	if err != nil {
		t.Fatalf("Error creating key from hex string")
	}
	encoded := Base58CheckEncode(key)
	expectedEncoded := "5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ"
	if encoded != expectedEncoded {
		t.Errorf("Base58CheckEncode should have returned '%s', but returned '%s' instead", expectedEncoded, encoded)
	}
}
