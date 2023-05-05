package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitub.com/sriramr98/fzmove/core"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Initializing dependencies")
    core.Init()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

