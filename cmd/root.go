/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitub.com/sriramr98/fzmove/core"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fzmove",
	Short: "Supercharge your terminal movements",
	Long: `Use fzmove to supercharge your terminal movements. 
         Open any project folder either in the terminal or in your favourite IDE/Editor.
         fzmover will map your project languages with your preferred applications so that opening your project in your own IDE or TextEditor will be buttery smooth.
  `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
    if !core.IsInitSuccess() {
      fmt.Println("You need run run `fzmove init` before you can use the application")
      return;
    }

    fmt.Println("Init Success .. Search to be Implemented");
  },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
    fmt.Fprintf(os.Stderr, "Oops!. There was an error '%s'", err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fzmove.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


