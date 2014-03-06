package mercer

import (
	"testing"
)

func TestBase58EncodingOfAnEmptySliceIsAnEmptyString(t *testing.T) {
	bytes :=  make([]byte, 0)
	encoded := Encode(bytes)
	if encoded != "" {
		t.Errorf("Expected empty string, got '%s'", encoded)
	}
}
