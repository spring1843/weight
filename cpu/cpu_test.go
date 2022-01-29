package cpu

import (
	"testing"
)

func TestRun(t *testing.T) {
	newCmd := NewCPUCommand()
	if !newCmd.Flags().HasFlags() {
		t.Fatal("expected command with flags but got no flags")
	}
	if err := ValidateFlags(); err != nil {
		t.Fatalf("failed validating default flags. Error : %s", err)
	}
	FlagDoNotChange = true

	breakRun = true
	if err := run("foo"); err == nil {
		t.Fatal("expected an error but got nil")
	}
	validateAndRun(nil, []string{})
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
