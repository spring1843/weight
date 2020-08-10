package main

import (
	"errors"
	"fmt"
	"os/exec"
)

func readLinuxCPULoad() (float32, error) {
	if topPath == "" {
		return 0, errors.New("empty top path")
	}

	cmd := exec.Command(topPath, []string{"-bn1"}...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return -1, fmt.Errorf("failed to execute the top command. Error : %w", err)
	}
	return parseLinuxTopOutput(string(out))
}

// parseLinuxTopOutput parses the output of top command in linux looking for CPU load
func parseLinuxTopOutput(topOutput string) (float32, error) {
	return parseTopOutput(&parseTopOutputParams{
		topOutput:         topOutput,
		keywordCPULoad:    "%Cpu(s):",
		keywordBeforeIdle: "ni,",
		keywordIdle:       " id",
	})
}
