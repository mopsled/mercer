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
