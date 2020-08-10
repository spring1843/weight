package main

import "testing"

func TestParseLinuxCPU(t *testing.T) {
	var tests = []struct {
		input  string
		output float32
		err    bool
	}{
		{"%Cpu(s):  0.0 us,  0.0 sy,  0.0 ni,60.0 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st", 40.00, false},
		{"%Cpu(s):  11.11 us,  22.22 sy,  0.0 ni,60.0 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st", 40.00, false},
		{"%Cpu(s):  11.11 us,  22.22 sy,  0.0 ni,6.0 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st", 94.00, false},
		{"%Cpu(s):  11.11 us,  22.22 sy,  0.0 ni,0.0 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st", 100.00, false},
		{"%Cpu(s):  11.11 us,  22.22 sy,  0.0 ni,1.5 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st", 98.50, false},
		{"%Cpu(s):  11.11 us,  22.22 sy,  0.0 ni, 1.5 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st", 98.50, false},
	}

	for i, test := range tests {
		got, err := parseLinuxTopOutput(test.input + "\n")
		if err != nil && !test.err {
			t.Fatalf("Failed test case #%d. Unexpected error. Error: %s", i, err)
		}
		if got != test.output {
			t.Fatalf("Failed test case #%d. Want %2.2f got %2.2f", i, test.output, got)
		}
	}
}
