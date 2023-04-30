package core

import (
	"fmt"

	"gitub.com/sriramr98/fzmove/utils"
)

func Init() {
  if !installFzF() {
    fmt.Println("======== INIT FAILED =========")
    return;
  }

  fmt.Println("===== INIT SUCCESS ======")
}

func installFzF() bool {
  if success := utils.CheckIfExists("fzf"); success {
    fmt.Println("FzF already installed")
    return true;
  }

  fmt.Println("Installing FzF...")
  if !isGitInstalled() {
    fmt.Println("Git Not Found")
    return false;
  }

  if !gitClone("https://github.com/junegunn/fzf.git", "~/.fzf") {
    fmt.Println("Unable to Clone FzF")
    return false;
  }

  if success := utils.RunScript("~/.fzf/install"); !success {
    fmt.Println("Unable to install FzF")
    return false;
  }

  return true;
}

func IsInitSuccess() bool {
  return false
}


