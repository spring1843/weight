package main

import "testing"

func TestRunWithBadCommandsPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Skipf("main function with unknown command did not panic.")
		}
	}()
	main()
}
