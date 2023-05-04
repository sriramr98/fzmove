package dependencies

import (
	"log"
	"os"
	"path"

	"gitub.com/sriramr98/fzmove/utils"
)

type FzF struct {
}

func (f FzF) IsAlreadyInstalled() bool {
  return utils.CheckIfExists("fzf")
}

func (f FzF) Install() bool {
  homeDirPath, err := os.UserHomeDir()
  if err != nil {
    log.Fatal("Home Dir not found ", err)
  }

  gitClient := utils.GitClient();

  if !gitClient.Clone("https://github.com/junegunn/fzf.git", path.Join(homeDirPath, ".fzf")) {
    return false;
  }

  if success := utils.RunScript(path.Join(homeDirPath, ".fzf", "install")); !success {
    return false;
  }

  return true;
}

func (f FzF) Name() string {
  return "fzf"
}

func NewFzFDependency() FzF {
  return FzF{}
}
