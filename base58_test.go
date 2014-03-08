package mercer

import (
	"encoding/hex"
	"testing"
)

func TestBase58EncodingOfAnEmptySliceIsAnEmptyString(t *testing.T) {
	bytes := make([]byte, 0)
	encoded := Encode(bytes)
	if encoded != "" {
		t.Errorf("Expected empty string, got '%s'", encoded)
	}
}

func TestBase58EncodingOf0x01Is2(t *testing.T) {
	bytes, err := hex.DecodeString("01")
	if err != nil {
		t.Fatalf("Error creating bytes from hex string")
	}
	encoded := Encode(bytes)
	if encoded != "2" {
		t.Errorf("Expected '2', got '%s'", encoded)
	}
}

func TestBase58EncodingOf0xffIs5Q(t *testing.T) {
	bytes, err := hex.DecodeString("ff")
	if err != nil {
		t.Fatalf("Error creating bytes from hex string")
	}
	encoded := Encode(bytes)
	if encoded != "5Q" {
		t.Errorf("Expected '5Q', got '%s'", encoded)
	}
}

func TestBase58EncodingOf0x0001Is12(t *testing.T) {
	bytes, err := hex.DecodeString("0001")
	if err != nil {
		t.Fatalf("Error creating bytes from hex string")
	}
	encoded := Encode(bytes)
	if encoded != "12" {
		t.Errorf("Expected '12', got '%s'", encoded)
	}
}

func TestBase58EncodingOfPrivateKeyIsCorrect(t *testing.T) {
	bytes, err := hex.DecodeString("800C28FCA386C7A227600B2FE50B7CAE11EC86D3BF1FBE471BE89827E19D72AA1D507A5B8D")
	if err != nil {
		t.Fatalf("Error creating bytes from hex string")
	}
	encoded := Encode(bytes)
	if encoded != "5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ" {
		t.Errorf("Expected '5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ', got '%s'", encoded)
	}
}

func TestBase58EncodingOfAnotherPrivateKeyIsCorrect(t *testing.T) {
	bytes, err := hex.DecodeString("ae2732ee016bdeaac6df3fe2e9b8a153689cc1c924e95fd563992387a5f14586")
	if err != nil {
		t.Fatalf("Error creating bytes from hex string")
	}
	encoded := Encode(bytes)
	if encoded != "CipcDUJfqM9pi4Dgq6W45ZLCnsfpGZsycEewdgGRnAW9" {
		t.Errorf("Expected 'CipcDUJfqM9pi4Dgq6W45ZLCnsfpGZsycEewdgGRnAW9', got '%s'", encoded)
	}
}
