package cpu

import (
	"testing"
)

func TestNewReader(t *testing.T) {
	if _, err := newReader("foo"); err == nil {
		t.Fatal("expected error but got none")
	}

	reader, err := newReader(osLinux)
	if err != nil {
		t.Fatalf("failed getting a new reader. Error: %s", err)
	}

	reader, err = newReader(osDarwin)
	if err != nil {
		t.Fatalf("failed getting a new reader. Error: %s", err)
	}

	if _, err := reader(); err != nil {
		t.Fatalf("failed getting a new reader. Error: %s", err)
	}

	oldTopCmd := topCmd
	topCmd = "topFooBar"
	if _, err := newReader("foo"); err == nil {
		t.Fatal("expected error but got none")
	}
	topCmd = oldTopCmd
}

func TestParser(t *testing.T) {
	if _, err := parseTopOutput(&parseTopOutputParams{"", "foo", "", ""}); err == nil {
		t.Fatal("expected error got nil")
	}
	if _, err := parseTopOutput(&parseTopOutputParams{"foobarbaz", "foo", "", ""}); err == nil {
		t.Fatal("expected error got nil")
	}
	if _, err := parseTopOutput(&parseTopOutputParams{"foobarbaz\n", "foo", "", ""}); err == nil {
		t.Fatal("expected error got nil")
	}
	if _, err := parseTopOutput(&parseTopOutputParams{"foobarbaz\n", "foo", "far", ""}); err == nil {
		t.Fatal("expected error got nil")
	}
	if _, err := parseTopOutput(&parseTopOutputParams{"foobarbaz\n", "foo", "bar", ""}); err == nil {
		t.Fatal("expected error got nil")
	}
	if _, err := parseTopOutput(&parseTopOutputParams{"foobarbaz\n", "foo", "", "far"}); err == nil {
		t.Fatal("expected error got nil")
	}
}
