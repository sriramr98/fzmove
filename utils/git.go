package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Git struct {}

func (g Git) IsInstalled() bool {

  cmd := exec.Command("git", "version")
  out, err := cmd.Output()

  if err != nil {
    log.Fatal(err)
    return false;
  }

  fmt.Printf("Git Version '%s' \n", string(out))
  return true;
}

func (g Git) Clone(repoUrl string, destPath string) bool {
  cmd := exec.Command("git" , "clone", repoUrl, destPath)
  cmd.Stdout = os.Stdout
  if err := cmd.Run(); err != nil {
    log.Fatalf("Unable to clone repo %s, error %s", repoUrl, err);
  }

  return true

}

var git *Git = nil;

func GitClient() Git {
  if git == nil {
    git = &Git{}
  }

  return *git
}
