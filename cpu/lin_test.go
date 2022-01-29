package cpu

import "testing"

func TestEitherFail(t *testing.T) {
	_, err1 := readOSXCPULoad()
	_, err2 := readLinuxCPULoad()

	if err1 != nil && err2 != nil {
		t.Fatalf("both readers fail, at least one must not fail. Err1: %s, Err2: %s", err1, err2)
	}
}

func TestFailOnNoTopPath(t *testing.T) {
	oldTopPath := topPath
	topPath = ""

	if _, err := readLinuxCPULoad(); err == nil {
		t.Fatal("Expected err but got nil ")
	}

	if _, err := readOSXCPULoad(); err == nil {
		t.Fatal("Expected err but got nil ")
	}
	topPath = oldTopPath
}

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
		{"%Cpu(s):  4.3 us, 13.0 sy,  0.0 ni, 82.5 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st", 17.50, false},
	}

	parseParams := linuxTopParams
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
