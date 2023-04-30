package core

import (
	"fmt"

	"gitub.com/sriramr98/fzmove/utils"
)

func Init() {
  if !isGitInstalled() {
    fmt.Println("======== INIT FAILED =========")
    return;
  }

  if !gitClone("https://github.com/junegunn/fzf.git", "~/.fzf") {
    fmt.Println("Unable to Clone FzF")
    return;
  }

  _, err := utils.RunScript("~/.fzf/install")

  if err != nil {
    fmt.Println("Unable to install FzF")
  }

  fmt.Println("===== INIT SUCCESS ======")
}

func IsInitSuccess() bool {
  return false
}


