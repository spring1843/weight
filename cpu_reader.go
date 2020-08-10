package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	topCmd         = "top"
	keyWordNewLine = "\n"
)

type (
	// cpuReader can read the CPU load using the OS API
	cpuReader func() (float32, error)

	// parseTopOutputParams holds the parameters passed in to parse a top command's output
	parseTopOutputParams struct {
		topOutput, keywordCPULoad, keywordBeforeIdle, keywordIdle string
	}
)

var topPath string

// newCPUReader returns a CPU reader for the current OS if it is supported
func newCPUReader() (cpuReader, error) {
	var err error
	topPath, err = loadTopPath()
	if err != nil {
		return nil, err
	}

	switch runtime.GOOS {
	case "darwin":
		return readOSXCPULoad, nil
	case "linux":
		return readLinuxCPULoad, nil
	}
	return nil, fmt.Errorf("OS %q is not currently supported. No method of finding CPU load", runtime.GOOS)
}

func loadTopPath() (string, error) {
	topPath, err := exec.LookPath(topCmd)
	if err != nil {
		return "", fmt.Errorf("Failed finding top path. Error : %s", err)
	}
	return topPath, nil
}

// parseTopOutput parses the cpu load out of the output of the top command in osx and linux
// The load is calculated by finding the idle percentage and then subtracting it from 100
// some load in both the system and usage percentages  hence it's easier just to parse the idle
// time and then use it to calculate the busy percentage
func parseTopOutput(params *parseTopOutputParams) (float32, error) {
	cpuIndex := strings.Index(params.topOutput, params.keywordCPULoad)
	if cpuIndex == -1 {
		return -1, fmt.Errorf("failed to parse CPU load. Keyword %q was not found in top command's output", params.keywordCPULoad)
	}
	params.topOutput = params.topOutput[cpuIndex+len(params.keywordCPULoad):]

	newLineIndex := strings.Index(params.topOutput, keyWordNewLine)
	if newLineIndex == -1 {
		return -1, fmt.Errorf("failed to parse CPU load. Keyword %q was not found in top command's output", keyWordNewLine)
	}
	params.topOutput = params.topOutput[:newLineIndex]
	cpuOutput := params.topOutput

	sysIndex := strings.Index(params.topOutput, params.keywordBeforeIdle)
	if cpuIndex == -1 {
		return -1, fmt.Errorf("failed to parse CPU load. Keyword %q was not found in top command's output", params.keywordBeforeIdle)
	}
	params.topOutput = params.topOutput[sysIndex+len(params.keywordBeforeIdle):]

	idleIndex := strings.Index(params.topOutput, params.keywordIdle)
	if idleIndex == -1 {
		return -1, fmt.Errorf("failed to parse CPU load. Keyword %q was not found in top command's output", params.keywordIdle)
	}
	params.topOutput = strings.TrimSpace(params.topOutput[:idleIndex])

	float, err := strconv.ParseFloat(params.topOutput, 32)
	if err != nil {
		return -1, fmt.Errorf("failed to parse float from %q taken out of %q", params.topOutput, cpuOutput)
	}
	return float32(100 - float), nil
}
