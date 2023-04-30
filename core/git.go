package core

import (
	"fmt"
	"log"
	"os/exec"
)

func isGitInstalled() bool {
  cmd := exec.Command("git", "version")
  out, err := cmd.Output()

  if err != nil {
    log.Fatal(err)
    return false;
  }

  fmt.Printf("Git Version '%s' \n", string(out))
  return true;
}

func gitClone(repoUrl string, destPath string) bool {
  cmd := exec.Command("git" , "clone", repoUrl, destPath)
  if _, err := cmd.Output(); err != nil {
    log.Fatalf("Unable to clone repo %s, error %s", repoUrl, err);
    return false
  }

  return true
}

