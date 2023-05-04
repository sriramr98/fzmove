package core

import (
	"fmt"
	"log"

	"gitub.com/sriramr98/fzmove/dependencies"
	"gitub.com/sriramr98/fzmove/utils"
)

func Init() {

  gitClient := utils.GitClient();

  if !gitClient.IsInstalled() {
    log.Fatal("Git Not Found. Please install git and rerun `fzmove init` to continue")
  }

  dependencies := dependencies.ALL_DEPENDENCIES

  for _, dep := range dependencies {
    name := dep.Name()
    if dep.IsAlreadyInstalled() {
      log.Printf("INFO: Dependency %s already installed\n", name)
      continue;
    }

    if success := dep.Install(); !success {
      log.Fatalf("ERROR: Unable to Install Dependency %s", name)
    }
  }

  fmt.Println("===== INIT SUCCESS ======")
}

func IsInitSuccess() bool {
  return true
}

