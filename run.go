package main

import (
	"fmt"
	"math"
	"time"
)

const (
	actionNothing = "Nothing"
	actionAdd     = "Add"
	actionRemove  = "Remove"
)

// cpuCheckSleepDuration is the amount of time we wait before each CPU read.
var cpuCheckDelayDuration time.Duration

// runCPU initiates the process of starting enough CPU loaders to meet the target
func runCPU() error {
	printf("Targeting %%%2.2f CPU load, With %s loader sleep duration and %d initial loader(s)", flagTargetCPULoad, loaderSleepDuration, flagInitialCPULoaders)
	addLoaders(flagInitialCPULoaders)

	cpuReader, err := newCPUReader()
	if err != nil {
		return fmt.Errorf("failed getting CPU reader. Error: %w", err)
	}

	return watchCPULoad(cpuReader)
}

// watchCPULoad keeps watching the CPU load, performs adjusting actions and prompts the outcome
func watchCPULoad(cpuReader cpuReader) error {
	for {
		load, err := cpuReader.read()
		if err != nil {
			return fmt.Errorf("failed getting CPU load. Error: %w", err)
		}

		loaderLen := loaderLen()
		action, count := actOnCPULoad(load, loaderLen)
		prompt(load, flagTargetCPULoad, fmt.Sprintf("%s %d", action, count), loaderLen)

		time.Sleep(cpuCheckDelayDuration)
	}
}

// actOnCPULoad intakes the current CPU load and the target, decides which one of
// the (nothing, add, remove) actions and how many of them are appropriate and then
// performs that action
func actOnCPULoad(load float32, loaderLen int) (string, int) {
	count := flagLoaderIncrements
	if count == 0 && flagLoaderIncrements != 0{
		count = int(math.Abs(float64(load - flagTargetCPULoad)))
	}

	action := actionNothing
	if load < flagTargetCPULoad {
		action = actionAdd
		addLoaders(count)
	}

	if load > flagTargetCPULoad && loaderLen > 0 {
		action = actionRemove
		if count > loaderLen {
			count = loaderLen
		}
		removeLoaders(count)
	}

	if action == actionNothing {
		count = 0
	}
	return action, count
}
