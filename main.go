package main

import (
	"time"

	"github.com/spf13/cobra"
)

var (
	flagQuiteMode             bool
	flagTargetCPULoad         float32
	flagLoaderSleepDuration   string
	flagCPUCheckDelayDuration string
	flagInitialCPULoaders     int
	flagLoaderIncrements      int
	flagDoNotChange           bool

	cmdRoot = &cobra.Command{
		Version: version,
		Use:     "weight",
	}
	cmdCPU = &cobra.Command{
		Use:   "cpu [options]",
		Short: "Put load on CPU",
		Long:  `Will attempt to reach a target CPU load by managing a number of CPU loaders that make it busy.`,
		Run: func(cmd *cobra.Command, args []string) {
			validateFlags()
			if err := runCPU(); err != nil {
				fatalf("Failed running CPU loader. Error: %s", err)
			}
		},
	}
)

func main() {
	registerFlags()
	cmdRoot.AddCommand(cmdCPU)
	if err := cmdRoot.Execute(); err != nil {
		fatalf("Failed executing command. Error: %s", err)
	}
}

func registerFlags() {
	cmdCPU.Flags().BoolVarP(&flagQuiteMode, "quiet", "q", false, "If set to true, informational logs will be silent")
	cmdCPU.Flags().Float32VarP(&flagTargetCPULoad, "target-cpu-load", "t", 50.0, "Float number indicating the desired CPU load. 50.00 indicated %50")
	cmdCPU.Flags().StringVarP(&flagLoaderSleepDuration, "loader-sleep-duration", "l", "100µs", "Duration of sleep before each attempt to busy the CPU. 0 would make each loader use 100% of the CPU and 10s would make each use almost none. Example values include 2ns, 2µs, 2ms, 2s, ...")
	cmdCPU.Flags().StringVarP(&flagCPUCheckDelayDuration, "cpu-check-delay-duration", "c", "2s", "Delay between each CPU load is read and corrective action is taken. Valid values include 2ns, 2µs, 2ms, 2s, ...")
	cmdCPU.Flags().IntVarP(&flagInitialCPULoaders, "initial-cpu-loaders", "i", 0, "Number of initial loaders that busy the CPU")
	cmdCPU.Flags().IntVarP(&flagLoaderIncrements, "loader-increments", "n", 0, "How many loaders to add or remove with each action, the default 0 value means use a smarter strategy rather than a fixed count")
	cmdCPU.Flags().BoolVarP(&flagDoNotChange, "do-not-change", "d", false, "If set to true then the checks for CPU read will be decided but no action will be taken")

}

func validateFlags() {
	var err error
	loaderSleepDuration, err = time.ParseDuration(flagLoaderSleepDuration)
	if err != nil {
		fatalf("Failed parsing busier-sleep-duration %q. Must be a duration value like 2ns, 2µs, 2ms, 10s, ...", flagLoaderSleepDuration)
	}
	if cpuCheckDelayDuration, err = time.ParseDuration(flagCPUCheckDelayDuration); err != nil {
		fatalf("Failed parsing busier-sleep-duration %q. Must be a duration value like 2ns, 2µs, 2ms, 10s, ...", cpuCheckDelayDuration)
	}
}

func prompt(current, target float32, action string, count int) {
	printf("Current: %%%2.2f, Target: %%%2.2f, Action: %s, Loader Count: %d", current, target, action, count)
}
