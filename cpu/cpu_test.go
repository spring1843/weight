package cpu

import "testing"

func TestRun(t *testing.T) {
	newCmd := NewCPUCommand()
	if !newCmd.Flags().HasFlags() {
		t.Fatal("Expected command with flags but got no flags")
	}
	if err := ValidateFlags(); err != nil {
		t.Fatalf("Failed validating default flags. Error : %s", err)
	}
	FlagDoNotChange = true

	breakRun = true
	if err := run(); err != nil {
		t.Fatalf("Failed to run. Error : %s", err)
	}
}

func TestValidateFlags(t *testing.T) {
	oldFlag := FlagLoaderSleepDuration
	FlagLoaderSleepDuration = "foo"
	if err := ValidateFlags(); err == nil {
		t.Fatal("Expected err but got none")
	}
	FlagLoaderSleepDuration = oldFlag

	oldFlag = FlagCPUCheckDelayDuration
	FlagCPUCheckDelayDuration = "foo"
	if err := ValidateFlags(); err == nil {
		t.Fatal("Expected err but got none")
	}
	FlagCPUCheckDelayDuration = oldFlag
}
