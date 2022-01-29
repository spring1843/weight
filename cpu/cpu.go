package cpu

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/spring1843/weight/log"
)

const (
	actionNothing = "Nothing"
	actionAdd     = "Add"
	actionRemove  = "Remove"
)

var (
	// The flags are each defined in NewCPUCommand

	// FlagTargetCPULoad is a flag defined for the CPU command
	FlagTargetCPULoad float32
	// FlagLoaderSleepDuration is a flag defined for the CPU command
	FlagLoaderSleepDuration string
	// FlagCPUCheckDelayDuration is a flag defined for the CPU command
	FlagCPUCheckDelayDuration string
	// FlagInitialCPULoaders is a flag defined for the CPU command
	FlagInitialCPULoaders int
	// FlagLoaderIncrements is a flag defined for the CPU command
	FlagLoaderIncrements int
	// FlagDoNotChange is a flag defined for the CPU command
	FlagDoNotChange bool

	// checkDelayDuration is the amount of time we wait before each CPU read.
	checkDelayDuration time.Duration

	cmd = &cobra.Command{
		Use:   "cpu [options]",
		Short: "Puts load on CPU",
		Long:  `Will attempt to reach a target CPU load by managing a number of CPU loaders that make it busy.`,
		Run:   validateAndRun,
	}

	runBreakMutex = new(sync.RWMutex)
	breakRun      = false
)

// NewCPUCommand returns a new CPU command
func NewCPUCommand() *cobra.Command {
	cmd.Flags().BoolVarP(&log.IsQuiteMode, "quite", "q", false, "No more informational logs")
	cmd.Flags().Float32VarP(&FlagTargetCPULoad, "target-cpu-load", "t", 50.0, "Float number indicating the desired CPU load. 50.00 indicated %50")
	cmd.Flags().StringVarP(&FlagLoaderSleepDuration, "loader-sleep-duration", "l", "100µs", "Duration of sleep before each attempt to busy the CPU. 0 would make each loader use 100% of the CPU and 10s would make each use almost none. Example values include 2ns, 2µs, 2ms, 2s, ...")
	cmd.Flags().StringVarP(&FlagCPUCheckDelayDuration, "cpu-check-delay-duration", "c", "2s", "Delay between each CPU load is read and corrective action is taken. Valid values include 2ns, 2µs, 2ms, 2s, ...")
	cmd.Flags().IntVarP(&FlagInitialCPULoaders, "initial-cpu-loaders", "i", 0, "Number of initial loaders that busy the CPU")
	cmd.Flags().IntVarP(&FlagLoaderIncrements, "loader-increments", "n", 0, "How many loaders to add or remove with each action, the default 0 value means use a smarter strategy rather than a fixed count")
	cmd.Flags().BoolVarP(&FlagDoNotChange, "do-not-change", "d", false, "If set to true then the checks for CPU read will be decided but no action will be taken")
	return cmd
}

func validateAndRun(cmd *cobra.Command, args []string) {
	if err := ValidateFlags(); err != nil {
		log.Fatalf("Failed running CPU loader. Error: %s", err)
	}
	if err := run(runtime.GOOS); err != nil {
		log.Fatalf("Failed running CPU loader. Error: %s", err)
	}
}

// run initiates the process of starting enough CPU loaders to meet the target
func run(os string) error {
	log.Printf("Targeting %%%2.2f CPU load, With %s loader sleep duration and %d initial loader(s)", FlagTargetCPULoad, loaderSleepDuration, FlagInitialCPULoaders)
	addLoaders(FlagInitialCPULoaders)

	cpuReader, err := newReader(os)
	if err != nil {
		return fmt.Errorf("failed getting CPU reader. Error: %w", err)
	}
	return watchLoad(cpuReader)
}

// ValidateFlags validates the runtime flags of CPU command
func ValidateFlags() error {
	var err error
	if loaderSleepDuration, err = time.ParseDuration(FlagLoaderSleepDuration); err != nil {
		return fmt.Errorf("failed parsing busier-sleep-duration %q. Must be a duration value like 2ns, 2µs, 2ms, 10s", FlagLoaderSleepDuration)
	}
	if checkDelayDuration, err = time.ParseDuration(FlagCPUCheckDelayDuration); err != nil {
		return fmt.Errorf("failed parsing busier-sleep-duration %q. Must be a duration value like 2ns, 2µs, 2ms, 10s", FlagCPUCheckDelayDuration)
	}
	return nil
}

// watchLoad keeps watching the CPU load, performs adjusting actions and prompts the outcome
func watchLoad(cpuReader reader) error {
	for {
		load, err := cpuReader()
		if err != nil {
			return fmt.Errorf("failed getting CPU load. Error: %w", err)
		}

		loaderLen := loaderLen()
		action, count := actOnCPULoad(load, loaderLen)
		prompt(load, FlagTargetCPULoad, fmt.Sprintf("%s %d", action, count), loaderLen)

		if FlagCPUCheckDelayDuration != zeroTime {
			time.Sleep(checkDelayDuration)
		}

		runBreakMutex.RLock()
		if breakRun {
			runBreakMutex.RUnlock()
			break
		}
		runBreakMutex.RUnlock()
	}
	return nil
}

// actOnCPULoad intakes the current CPU load and the target, decides which one of
// the (nothing, add, remove) actions and how many of them are appropriate and then
// performs that action
func actOnCPULoad(load float32, loaderLen int) (string, int) {
	count := FlagLoaderIncrements
	if count == 0 && !FlagDoNotChange {
		count = int(math.Abs(float64(load - FlagTargetCPULoad)))
	}

	action := actionNothing
	if load < FlagTargetCPULoad {
		action = actionAdd
		addLoaders(count)
	}

	if load > FlagTargetCPULoad && loaderLen > 0 {
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

func prompt(current, target float32, action string, count int) {
	log.Printf("Current: %%%2.2f, Target: %%%2.2f, Action: %s, Loader Count: %d", current, target, action, count)
}
