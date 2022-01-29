package log

import "log"

// IsQuiteMode indicates if we are in quite mode or not
var IsQuiteMode bool

// Printf calls the standard Printf logger unless we are in quite mode
func Printf(format string, a ...interface{}) {
	if IsQuiteMode {
		return
	}
	log.Printf(format, a...)
}

// Fatalf makes a call to Printf and then os.Exit(1)
func Fatalf(format string, a ...interface{}) {
	log.Fatalf(format, a...)
}
