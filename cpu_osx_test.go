package main

import "testing"

func TestParseOSXCPU(t *testing.T) {
	var tests = []struct {
		input  string
		output float32
		err    bool
	}{
		{"CPU usage: 0.0% user, 33.33% sys, 60.00% idle", 40.00, false},
		{"CPU usage: 11.11% user, 22.22% sys, 60.00% idle", 40.00, false},
		{"CPU usage: 11.11% user, 22.22% sys, 6.00% idle", 94.00, false},
		{"CPU usage: 11.11% user, 22.22% sys, 0.00% idle", 100.00, false},
		{"CPU usage: 11.11% user, 22.22% sys, 1.5% idle", 98.50, false},
	}

	parseParams := osxTopParams
	for i, test := range tests {
		parseParams.topOutput = test.input + "\n"
		got, err := parseTopOutput(parseParams)
		if err != nil && !test.err {
			t.Fatalf("Failed test case #%d. Unexpected error. Error: %s", i, err)
		}
		if got != test.output {
			t.Fatalf("Failed test case #%d. Want %2.2f got %2.2f", i, test.output, got)
		}
	}
}
