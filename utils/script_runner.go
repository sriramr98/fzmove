package utils

import "os/exec"

func RunScript(scriptPath string) (string, error) {
  cmd := exec.Command("sh", scriptPath)
  out, err := cmd.Output()

  return string(out), err
}
