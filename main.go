package main

import (
	"github.com/spring1843/weight/cpu"
	"github.com/spring1843/weight/log"

	"github.com/spf13/cobra"
)

var cmdRoot = &cobra.Command{
	Version: version,
	Use:     "weight",
}

func main() {
	cmdRoot.AddCommand(cpu.NewCPUCommand())
	if err := cmdRoot.Execute(); err != nil {
		log.Fatalf("Failed executing command. Error: %s", err)
	}
}
