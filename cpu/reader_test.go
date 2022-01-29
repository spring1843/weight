package cpu

import "testing"

func TestNewReader(t *testing.T) {
	reader, err := newReader()
	if err != nil {
		t.Fatalf("failed getting a new reader. Error: %s", err)
	}

	if _, err := reader(); err != nil {
		t.Fatalf("failed getting a new reader. Error: %s", err)
	}
}
