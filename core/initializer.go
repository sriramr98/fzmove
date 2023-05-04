package core

import (
	"fmt"
	"log"
	"os"
	"path"

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

  if success := installFzf(); !success {
    log.Fatal("Unable to install FzF");
    return false;
  }
  
  return true;
}

func IsInitSuccess() bool {
  return true
}

func installFzf() bool {
  homeDirPath, err := os.UserHomeDir()
  if err != nil {
    log.Fatal("Home Dir not found ", err)
  }

  if !gitClone("https://github.com/junegunn/fzf.git", path.Join(homeDirPath, ".fzf")) {
    fmt.Println("Unable to Clone FzF")
    return false;
  }

  if success := utils.RunScript(path.Join(homeDirPath, ".fzf", "install")); !success {
    fmt.Println("Unable to install FzF")
    return false;
  }

  return true;
}

