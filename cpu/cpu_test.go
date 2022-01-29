package cpu

import (
	"errors"
	"fmt"
	"testing"

	"github.com/spring1843/weight/log"
)

var (
	calledLogger      = false
	mockedReaderErr   error
	mockedReaderValue float32
)

func mockedLogger(format string, v ...interface{}) {
	calledLogger = true
}

func mockedReader() (float32, error) {
	return mockedReaderValue, mockedReaderErr
}

func TestRun(t *testing.T) {
	newCmd := NewCPUCommand()
	if !newCmd.Flags().HasFlags() {
		t.Fatal("expected command with flags but got no flags")
	}
	if err := ValidateFlags(); err != nil {
		t.Fatalf("failed validating default flags. Error : %s", err)
	}
	oldFlagDoNotChange := FlagDoNotChange
	FlagDoNotChange = true

	breakRun = true
	if err := run("foo"); err == nil {
		t.Fatal("expected an error but got nil")
	}

	oldFlagLoaderSleepDuration := FlagLoaderSleepDuration
	FlagLoaderSleepDuration = "foo"
	log.FatalfFunc = mockedLogger
	validateAndRun(nil, []string{})
	if !calledLogger {
		t.Fatal("invalid duration did not trigger fatal")
	}
	FlagLoaderSleepDuration = oldFlagLoaderSleepDuration
	calledLogger = false

	oldFlagCPUCheckDelayDuration := FlagCPUCheckDelayDuration
	FlagCPUCheckDelayDuration = "foo"
	log.FatalfFunc = mockedLogger
	validateAndRun(nil, []string{})
	if !calledLogger {
		t.Fatal("invalid duration did not trigger fatal")
	}
	FlagCPUCheckDelayDuration = oldFlagCPUCheckDelayDuration
	calledLogger = false

	oldRuntimeOperatingSystem := runtimeOperatingSystem
	runtimeOperatingSystem = "foo"
	log.FatalfFunc = mockedLogger
	validateAndRun(nil, []string{})
	if !calledLogger {
		t.Fatal("invalid os did not trigger fatal")
	}
	runtimeOperatingSystem = oldRuntimeOperatingSystem

	validateAndRun(nil, []string{})
	FlagDoNotChange = oldFlagDoNotChange
}

func TestValidateFlags(t *testing.T) {
	oldFlag := FlagLoaderSleepDuration
	FlagLoaderSleepDuration = "foo"
	if err := ValidateFlags(); err == nil {
		t.Fatal("expected err but got none")
	}
	FlagLoaderSleepDuration = oldFlag

	oldFlag = FlagCPUCheckDelayDuration
	FlagCPUCheckDelayDuration = "foo"
	if err := ValidateFlags(); err == nil {
		t.Fatal("expected err but got none")
	}
	FlagCPUCheckDelayDuration = oldFlag
}

func TestWatchLoadWithMockedReader(t *testing.T) {
	mockedReaderErr = errors.New("foo")
	if err := watchLoad(mockedReader); err == nil {
		t.Fatal("expected error but got nil")
	}

	mockedReaderErr = nil
	breakRun = false
	watch := func() {
		if err := watchLoad(mockedReader); err != nil {
			panic(fmt.Errorf("mocked reader exited with error. Error :%s", err))
		}
	}
	go watch()
	runBreakMutex.Lock()
	breakRun = true
	runBreakMutex.Unlock()

	emptyLoaders()
}

func TestActOnCPULoad(t *testing.T) {
	emptyLoaders()
	oldFlagTargetCPULoad := FlagTargetCPULoad
	oldFlagDoNotChange := FlagDoNotChange
	FlagDoNotChange = false
	FlagTargetCPULoad = 3
	action, count := actOnCPULoad(1.00, 1)
	if action != actionAdd {
		t.Fatalf("did not get expected action. Want %s, got %s.", actionNothing, action)
	}
	if count != 2 {
		t.Fatalf("did not get expected count. Want %d, got %d.", 2, count)
	}
	FlagDoNotChange = oldFlagDoNotChange

	emptyLoaders()
	FlagTargetCPULoad = 0
	action, count = actOnCPULoad(0.00, 1)
	if action != actionNothing {
		t.Fatalf("did not get expected action. Want %s, got %s.", actionNothing, action)
	}
	if count != 0 {
		t.Fatalf("did not get expected count. Want %d, got %d.", 0, count)
	}
	FlagTargetCPULoad = oldFlagTargetCPULoad

	emptyLoaders()
	FlagTargetCPULoad = 0
	FlagLoaderIncrements = 5
	addLoaders(3)
	action, count = actOnCPULoad(2.00, 3)
	if action != actionRemove {
		t.Fatalf("did not get expected action. Want %s, got %s.", actionNothing, action)
	}
	if count != 3 {
		t.Fatalf("did not get expected count. Want %d, got %d.", 3, count)
	}
	FlagTargetCPULoad = oldFlagTargetCPULoad
	emptyLoaders()
}

func emptyLoaders() {
	removeLoaders(loaderLen())
}
