package log

import (
	"testing"
)

var calledLogger = false

func mockedLogger(format string, v ...interface{}) {
	calledLogger = true
}

func TestQuiteMode(t *testing.T) {
	printfFunc = mockedLogger
	FatalfFunc = mockedLogger
	Printf("foo")
	if !calledLogger {
		t.Fatal("did not print in normal mode")
	}

	calledLogger = false
	Fatalf("foo")
	if !calledLogger {
		t.Fatal("did not fatal in normal mode")
	}

	calledLogger = false
	IsQuiteMode = true
	Printf("foo")
	if calledLogger {
		t.Fatal("printed in quite mode")
	}

	calledLogger = false
	IsQuiteMode = true
	Fatalf("foo")
	if !calledLogger {
		t.Fatal("did not call fatal in quite mode")
	}
}
