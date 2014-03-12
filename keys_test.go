package mercer

import (
	"bytes"
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
