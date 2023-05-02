/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"gitub.com/sriramr98/fzmove/cmd"
	"gitub.com/sriramr98/fzmove/config"
)

func main() {
  config.InitConfig()
	cmd.Execute()
}
