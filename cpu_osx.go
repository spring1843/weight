package main

import (
	"errors"
	"fmt"
	"os/exec"
)

// osxReader can read CPU load on OSX
type osxReader struct {
	path string
}

func readOSXCPULoad() (float32, error) {
	if topPath == "" {
		return 0, errors.New("empty top path")
	}

	cmd := exec.Command(topPath, []string{"-l", "1", "-n", "0"}...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return -1, fmt.Errorf("failed to execute the top command. Error : %w", err)
	}
	return parseOSXTopOutput(string(out))
}

// parseOSXTopOutput parses the output of top command in OSX looking for CPU load
func parseOSXTopOutput(topOutput string) (float32, error) {
	return parseTopOutput(&parseTopOutputParams{
		topOutput:         topOutput,
		keywordCPULoad:    "CPU usage:",
		keywordBeforeIdle: "sys, ",
		keywordIdle:       "% idle",
	})
}
