package main

import "log"

func printf(format string, a ...interface{}) {
	if flagQuiteMode {
		return
	}
	log.Printf(format, a...)
}

func fatalf(format string, a ...interface{}) {
	log.Fatalf(format, a...)
}
