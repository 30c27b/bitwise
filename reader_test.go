package bitwise_test

import (
	"bytes"
	"testing"

	"github.com/30c27b/bitwise"
)

func TestReader(t *testing.T) {
	buf := []byte{0b10100101, 0b11101110}

	r := bitwise.NewReader(bytes.NewReader(buf))

	n, err := r.ReadBits(7)
	if err != nil {
		t.Fatal("Unexpected failure when calling ReadBits:", err)
	}
	if n != 0b1010010 {
		t.Errorf("Expected %b, got %b", 0b1010010, n)
	}

	n, err = r.ReadBits(1)
	if err != nil {
		t.Fatal("Unexpected failure when calling ReadBits:", err)
	}
	if n != 0b1 {
		t.Errorf("Expected %b, got %b", 0b1, n)
	}

	n, err = r.ReadBits(4)
	if err != nil {
		t.Fatal("Unexpected failure when calling ReadBits:", err)
	}
	if n != 0b1110 {
		t.Errorf("Expected %b, got %b", 0b1110, n)
	}

	n, err = r.ReadBits(4)
	if err != nil {
		t.Fatal("Unexpected failure when calling ReadBits:", err)
	}
	if n != 0b1110 {
		t.Errorf("Expected %b, got %b", 0b1110, n)
	}

	_, err = r.ReadBits(1)
	if err == nil {
		t.Fatal("Expected EOF error, got nil")
	}
}
