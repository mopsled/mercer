package mercer

import (
	"bytes"
	"encoding/hex"
	"reflect"
	"testing"
)

func TestPrivateKeyIsAByteSlice(t *testing.T) {
	key, err := PrivateKey()
	if err != nil {
		t.Fatalf("PrivateKey() returned an error: %s", err)
	}
	kind := reflect.TypeOf(key).Elem().Kind()
	if kind != reflect.Uint8 {
		t.Errorf("A private key should be a byte slice")
	}
}

func TestPrivateKeyIsTheRightLength(t *testing.T) {
	key, err := PrivateKey()
	if err != nil {
		t.Fatalf("PrivateKey() returned an error: %s", err)
	}
	if len(key) != 64 {
		t.Errorf("A private key should be a 64-byte slice")
	}
}

func TestPrivateKeysAreGeneratedDifferent(t *testing.T) {
	key1, err := PrivateKey()
	if err != nil {
		t.Fatalf("PrivateKey() returned an error: %s", err)
	}
	key2, err := PrivateKey()
	if err != nil {
		t.Fatalf("PrivateKey() returned an error: %s", err)
	}
	if bytes.Equal(key1, key2) {
		t.Errorf("Two generated private keys should have different bytes")
	}
}

func TestChecksumBytes(t *testing.T) {
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
