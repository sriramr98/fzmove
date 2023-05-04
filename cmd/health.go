package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
  Use: "health",
  Short: "Check the health of the CLI..",
  Long: "Check the health of the CLI to ensure all dependencies are properly installed",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Checking Health");
  },
}
