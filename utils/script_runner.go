package utils

import "os/exec"

func RunScript(scriptPath string) bool {
  return RunShellCommand("sh", scriptPath)
}

func RunShellCommand(name string, params ...string) bool {
  cmd := exec.Command(name, params...)
  if _, err := cmd.Output(); err != nil {
    return false
  }
  return true
}

func CheckIfExists(program string) bool {
  return RunShellCommand("command", "-v", program)
}
