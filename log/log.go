package log

import "log"

type logFunc func(format string, v ...interface{})

var (
	// IsQuiteMode indicates if we are in quite mode or not
	IsQuiteMode bool
	printfFunc  logFunc = log.Printf
	// FatalfFunc is the function that will be called when Fatalf is called
	FatalfFunc logFunc = log.Fatalf
)

// Printf calls the standard Printf logger unless we are in quite mode
func Printf(format string, a ...interface{}) {
	if IsQuiteMode {
		return
	}
	printfFunc(format, a...)
}

// Fatalf makes a call to Printf and then os.Exit(1)
func Fatalf(format string, a ...interface{}) {
	FatalfFunc(format, a...)
}
