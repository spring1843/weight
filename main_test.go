package main

import (
	"os"
	"testing"
)

func TestRunWithBadCommandsPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("main function with unknown command did not panic.")
		}
	}()
	os.Args = []string{"weight", "foo"}
	main()
}
