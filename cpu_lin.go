package main

import (
	"errors"
)

var (
	linuxTopArgs   = []string{"-bn1"}
	linuxTopParams = &parseTopOutputParams{
		keywordCPULoad:    "%Cpu(s):",
		keywordBeforeIdle: "ni,",
		keywordIdle:       " id",
	}
)

func readLinuxCPULoad() (float32, error) {
	if topPath == "" {
		return 0, errors.New("empty top path")
	}
	return runTopAndParse(linuxTopArgs, linuxTopParams)
}
