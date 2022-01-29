package cpu

import (
	"errors"
)

var (
	linuxTopArgs   = []string{"-bn1"}
	linuxTopParams = &parseTopOutputParams{
		cpuLoad:    "%Cpu(s):",
		beforeIdle: "ni,",
		idle:       " id",
	}
)

func readLinuxCPULoad() (float32, error) {
	if topPath == "" {
		return 0, errors.New("empty top path")
	}
	return runTopAndParse(linuxTopArgs, linuxTopParams)
}
