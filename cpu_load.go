package main

import (
	"sync"
	"time"
)

// loader is able to put load on CPU
type loader struct {
	stopped bool
	lock    *sync.RWMutex
}

// zeroTime is used as a convention from user input to mean zero duration
const zeroTime = "0"

var (
	// loaders is all the currently active loaders
	loaders []*loader

	// loaderSleepDuration is a duration that loaders sleep after each execution.
	// 0ns would make each routine take all the available CPU, 10s would make each
	// routine take almost nothing. This value can be used to get more accurate results
	// on different machines and in different conditions.
	loaderSleepDuration time.Duration
)

func newLoader() *loader {
	return &loader{
		stopped: false,
		lock:    new(sync.RWMutex),
	}
}

// start puts some load on the CPU
// the amount of load has an inverse effect from loaderSleepDuration
func (l *loader) start() {
	for {
		l.lock.RLock()
		isStopped := l.stopped
		l.lock.RUnlock()

		if isStopped {
			break
		}

		time.Sleep(loaderSleepDuration)
	}
}

// startAndDoNotSleep is like start but does not sleep
// used to avoid the extra checks
func (l *loader) startAndDoNotSleep() {
	for {
		continue
	}
}

func (l *loader) end() {
	l.lock.Lock()
	l.stopped = true
	l.lock.Unlock()
}

// addLoaders adds n new loaders to the currently running loaders and starts them
func addLoaders(n int) {
	for i := 0; i < n; i++ {
		loaders = append(loaders, newLoader())

		if flagLoaderSleepDuration == zeroTime && !flagDoNotChange {
			go loaders[len(loaders)-1].startAndDoNotSleep()
			continue
		}
		go loaders[len(loaders)-1].start()
	}
}

// removeLoaders ends n loaders and removes them
func removeLoaders(n int) {
	for i := 0; i < n; i++ {
		loaders[0].end()
		loaders = loaders[1:]
	}
}

// loaderLen returns the number of currently active loaders
func loaderLen() int {
	return len(loaders)
}
