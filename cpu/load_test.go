package cpu

import (
	"testing"
	"time"
)

func TestLoader(t *testing.T) {
	loader := newLoader()
	loader.stopped = false
	go loader.start()
	time.Sleep(loaderSleepDuration)
	addLoaders(2)
	time.Sleep(loaderSleepDuration)
	want := 2
	if got := loaderLen(); got != want {
		t.Fatalf("Did not get the expected loader count. Want %d, got %d", want, got)
	}
	removeLoaders(2)
	want = 0
	if got := loaderLen(); got != want {
		t.Fatalf("Did not get the expected loader count. Want %d, got %d", want, got)
	}
	loader.end()
}

func TestLoaderNoChange(t *testing.T) {
	oldFlagLoaderSleepDuration := FlagLoaderSleepDuration
	oldFlagDoNotChange := FlagDoNotChange
	FlagDoNotChange = false
	FlagLoaderSleepDuration = zeroTime
	loader := newLoader()
	loader.stopped = false
	go loader.start()
	time.Sleep(loaderSleepDuration)
	addLoaders(2)
	time.Sleep(loaderSleepDuration)
	want := 2
	if got := loaderLen(); got != want {
		t.Fatalf("Did not get the expected loader count. Want %d, got %d", want, got)
	}
	loader.end()
	FlagDoNotChange = oldFlagDoNotChange
	FlagLoaderSleepDuration = oldFlagLoaderSleepDuration

	loader.lock.Lock()
	loader.stopped = false
	loader.lock.Unlock()

	go loader.start()
	loader.lock.Lock()
	loader.stopped = true
	loader.lock.Unlock()
	loader.startAndDoNotSleep()
}
