package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:     "hexa",
	Short:   "hexagonal arichtecture",
	Version: "1.0.0",
}

// Execute the root command using cobra CLI generator
func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
