package main

import (
	"errors"
)

var (
	osxTopArgs   = []string{"-l", "1", "-n", "0"}
	osxTopParams = &parseTopOutputParams{
		cpuLoad:    "CPU usage:",
		beforeIdle: "sys, ",
		idle:       "% idle",
	}
)

func readOSXCPULoad() (float32, error) {
	if topPath == "" {
		return 0, errors.New("empty top path")
	}
	return runTopAndParse(osxTopArgs, osxTopParams)
}
