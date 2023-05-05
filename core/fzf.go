package core

import "gitub.com/sriramr98/fzmove/utils"

func ProcessFzF(input string) (string, error) {
  return utils.RunShellCommand("echo", input , "|", "fzf")
}
