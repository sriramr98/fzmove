package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gitub.com/sriramr98/fzmove/dependencies"
)

var healthCmd = &cobra.Command{
  Use: "health",
  Short: "Check the health of the CLI..",
  Long: "Check the health of the CLI to ensure all dependencies are properly installed",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Checking Health");

    dependencies := dependencies.ALL_DEPENDENCIES

    success_deps, failed_deps := []string{}, []string{}
    for _, dep := range dependencies {
      name := dep.Name()
      if !dep.IsInstalled() {
        if ! silent {
          log.Printf("ERROR: Dependency %s is not installed", name)
        }
        failed_deps = append(failed_deps, name)
      } else {
        if !silent {
          log.Printf("INFO: Dependency %s installed", name)
        }
        success_deps = append(success_deps, name)
      }
    }

    if silent {
      log.Println("INFO: == SUCCESS ==")
      for _, dep := range success_deps {
        log.Println(dep)
      }

      log.Println("INFO: == FAILURE ==")
      for _, dep := range failed_deps {
        log.Println(dep)
      }
    }


  },
}

var silent bool

func init() {

  healthCmd.Flags().BoolVarP(&silent, "silent", "s", false, "silent output - only print summary")

  rootCmd.AddCommand(healthCmd)
}
