package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spring1843/weight/cpu"
)

var cmdRoot = &cobra.Command{
	Version: version,
	Use:     "weight",
}

func main() {
	cmdRoot.AddCommand(cpu.NewCPUCommand())
	if err := cmdRoot.Execute(); err != nil {
		panic(fmt.Errorf("failed executing command. Error: %s", err))
	}
}
