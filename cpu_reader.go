package main

import (
	"fmt"
	"runtime"
)

// cpuReader can read the CPU load using the OS API
type cpuReader interface {
	read() (float32, error)
}

// newCPUReader returns a CPU reader for the current OS if it is supported
func newCPUReader() (cpuReader, error) {
	switch runtime.GOOS {
	case "darwin":
		return newOSXReader()
	case "linux":
		return newLinuxReader()
	}
	return nil, fmt.Errorf("OS %q is not currently supported. No method of finding CPU load", runtime.GOOS)
}
