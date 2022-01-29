package cpu

import (
	"runtime"
	"testing"
)

func TestNewReader(t *testing.T) {
	if _, err := newReader("foo"); err == nil {
		t.Fatal("expected error but got none")
	}

	if _, err := newReader(osLinux); err != nil {
		t.Fatalf("failed getting a new reader. Error: %s", err)
	}

	if _, err := newReader(osDarwin); err != nil {
		t.Fatalf("failed getting a new reader. Error: %s", err)
	}

	if _, err := newReader(runtime.GOOS); err != nil {
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
