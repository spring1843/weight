package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// osxReader can read CPU load on OSX
type osxReader struct {
	path string
}

const (
	topCmd         = "top"
	keyWordNewLine = "\n"
)

func newOSXReader() (*osxReader, error) {
	path, err := exec.LookPath(topCmd)
	if err != nil {
		return nil, fmt.Errorf("failed finding the location of the top command. Error : %w", err)
	}

	return &osxReader{
		path: path,
	}, nil
}

func (o *osxReader) read() (float32, error) {
	cmd := exec.Command(o.path, []string{"-l", "1", "-n", "0"}...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return -1, fmt.Errorf("failed to execute the top command. Error : %w", err)
	}
	return parseOSXTopOutput(string(out))
}

// parseOSXTopOutput parses the output of top command looking for CPU load
// The load is calculated by finding the idle percentage and then subtracting it from 100
// This is the same as adding User + System percentages, it appears that running this app puts
// some load in both the system and usage percentages  hence it's easier just to parse the idle
// time and then use it to calculate the busy percentage
func parseOSXTopOutput(topOutput string) (float32, error) {
	const (
		keywordCPULoad = "CPU usage:"
		keywordSys     = "sys, "
		keywordIdle    = "% idle"
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

	sysIndex := strings.Index(topOutput, keywordSys)
	if cpuIndex == -1 {
		return -1, fmt.Errorf("failed to parse CPU load. Keyword %q was not found in top command's output", keywordSys)
	}
	topOutput = topOutput[sysIndex+len(keywordSys):]

	idleIndex := strings.Index(topOutput, keywordIdle)
	if idleIndex == -1 {
		return -1, fmt.Errorf("failed to parse CPU load. Keyword %q was not found in top command's output", keywordIdle)
	}
	topOutput = topOutput[:idleIndex]

	float, err := strconv.ParseFloat(topOutput, 32)
	if err != nil {
		return -1, fmt.Errorf("failed to parse float from %q taken out of %q", topOutput, cpuOutput)
	}
	return float32(100 - float), nil
}
