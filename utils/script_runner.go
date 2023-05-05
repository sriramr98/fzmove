package utils

import "os/exec"

func RunScript(scriptPath string) bool {
  _, err := RunShellCommand("sh", scriptPath)
  if err != nil {
    return false
  } 
  return true
}

func RunShellCommand(name string, params ...string) (string, error) {
  cmd := exec.Command(name, params...)
  res, err := cmd.Output();   

  if err != nil {
    return "", err
  }

  return string(res), err
}

func CheckIfExists(program string) bool {
  _, err := RunShellCommand("command", "-v", program)
  if err != nil {
    return false
  }
  return true
}
