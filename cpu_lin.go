package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// linuxReader can read CPU load on Linux
type linuxReader struct {
	path string
}

func newLinuxReader() (*linuxReader, error) {
	path, err := exec.LookPath(topCmd)
	if err != nil {
		return nil, fmt.Errorf("failed finding the location of the top command. Error : %w", err)
	}

	return &linuxReader{
		path: path,
	}, nil
}

func (o *linuxReader) read() (float32, error) {
	cmd := exec.Command(o.path, []string{"-bn1"}...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return -1, fmt.Errorf("failed to execute the top command. Error : %w", err)
	}
	return parseLinuxTopOutput(string(out))
}

// parseLinuxTopOutput parses the output of top command looking for CPU load
// The load is calculated by finding the idle percentage and then subtracting it from 100
func parseLinuxTopOutput(topOutput string) (float32, error) {
	const (
		keywordCPULoad = "%Cpu(s):"
		keywordNi      = "ni,"
		keywordIdle    = " id"
	)

	cpuIndex := strings.Index(topOutput, keywordCPULoad)
	if cpuIndex == -1 {
		return -1, fmt.Errorf("failed to parse CPU load. Keyword %q was not found in top command's output", keywordCPULoad)
	}
	topOutput = topOutput[cpuIndex+len(keywordCPULoad):]

	newLineIndex := strings.Index(topOutput, keyWordNewLine)
	if newLineIndex == -1 {
		return -1, fmt.Errorf("failed to parse CPU load. Keyword %q was not found in top command's output", keyWordNewLine)
	}
	topOutput = topOutput[:newLineIndex]
	cpuOutput := topOutput

	sysIndex := strings.Index(topOutput, keywordNi)
	if cpuIndex == -1 {
		return -1, fmt.Errorf("failed to parse CPU load. Keyword %q was not found in top command's output", keywordNi)
	}
	topOutput = topOutput[sysIndex+len(keywordNi):]

	idleIndex := strings.Index(topOutput, keywordIdle)
	if idleIndex == -1 {
		return -1, fmt.Errorf("failed to parse CPU load. Keyword %q was not found in top command's output", keywordIdle)
	}
	topOutput = strings.TrimSpace(topOutput[:idleIndex])

	float, err := strconv.ParseFloat(topOutput, 32)
	if err != nil {
		return -1, fmt.Errorf("failed to parse float from %q taken out of %q", topOutput, cpuOutput)
	}
	return float32(100 - float), nil
}
