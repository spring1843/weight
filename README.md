# weight

weight is multi OS command line application that tries to set the CPU load of the machine to the given value.

The name comes from the idea that it puts a fake load on the CPU. The project aims to provide similar capabilities for other system resources.

## Examples

`weight cpu -t 90.00` aims to set the CPU to 90% by running fake loads in a strategic manner.

```ANSI
2020/08/07 22:18:56 Targeting %90.00 CPU load, With 100µs loader sleep duration and 0 initial loader(s)
2020/08/07 22:18:57 Current: %34.54, Target: %90.00, Action: Add 55, Loader Count: 0
2020/08/07 22:18:59 Current: %52.71, Target: %90.00, Action: Add 37, Loader Count: 55
2020/08/07 22:19:02 Current: %60.55, Target: %90.00, Action: Add 29, Loader Count: 92
2020/08/07 22:19:04 Current: %62.90, Target: %90.00, Action: Add 27, Loader Count: 121
2020/08/07 22:19:07 Current: %68.31, Target: %90.00, Action: Add 21, Loader Count: 148
2020/08/07 22:19:09 Current: %66.22, Target: %90.00, Action: Add 23, Loader Count: 169
2020/08/07 22:19:12 Current: %71.62, Target: %90.00, Action: Add 18, Loader Count: 192
2020/08/07 22:19:14 Current: %73.08, Target: %90.00, Action: Add 16, Loader Count: 210
2020/08/07 22:19:17 Current: %72.10, Target: %90.00, Action: Add 17, Loader Count: 226
2020/08/07 22:19:19 Current: %96.70, Target: %90.00, Action: Remove 6, Loader Count: 243
2020/08/07 22:19:22 Current: %74.13, Target: %90.00, Action: Add 15, Loader Count: 237
```

See all available options by running `weight cpu -h`

```ANSI
Will attempt to reach a target CPU load by managing a number of CPU loaders that make it busy.

Usage:
  weight cpu [options] [flags]

Flags:
 Will attempt to reach a target CPU load by managing a number of CPU loaders that make it busy.

Usage:
  weight cpu [options] [flags]

Flags:
  -c, --cpu-check-delay-duration string   Delay between each CPU load is read and corrective action is taken. Valid values include 2ns, 2µs, 2ms, 2s, ... (default "2s")
  -d, --do-not-change                     If set to true then the checks for CPU read will be decided but no action will be taken
  -h, --help                              help for cpu
  -i, --initial-cpu-loaders int           Number of initial loaders that busy the CPU
  -n, --loader-increments int             How many loaders to add or remove with each action, the default 0 value means use a smarter strategy rather than a fixed count
  -l, --loader-sleep-duration string      Duration of sleep before each attempt to busy the CPU. 0 would make each loader use 100% of the CPU and 10s would make each use almost none. Example values include 2ns, 2µs, 2ms, 2s, ... (default "100µs")
  -q, --quiet                             If set to true, informational logs will be silent
  -t, --target-cpu-load float32           Float number indicating the desired CPU load. 50.00 indicated %50 (default 50)
```

## Installation

### Build from source

1. Download and install [Go](https://golang.org/dl/). Currently tested Go version in (1.14+).
2. Run `go get -u github.com/spring1843/weight` to download the source file locally
3. Run `make build`
