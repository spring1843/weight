# weight

weight is a highly customizable multi OS command line application that puts enough load on the CPU to meet the user target.

[![Build Status](https://github.com/spring1843/weight/workflows/PR/badge.svg)](https://github.com/spring1843/weight/actions)
[![Build Status](https://goreportcard.com/badge/github.com/spring1843/weight)](https://goreportcard.com/report/github.com/spring1843/weight)

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

## Installation

### Docker

1- Install Docker
2- Run `docker build -t spring1843/weight .` in the project directory
3- Run `docker run spring1843/weight:latest cpu`

### Build from source

1. Download and install [Go](https://golang.org/dl/) v1.14+
2. Run `go get -u github.com/spring1843/weight` and you should be able to run the `weight` command

If this fails because of environmental issues you can try running `make build` in the source directory to build the weight binary file.
