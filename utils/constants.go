package utils
// Actions
const ACTION_INIT = "init"
const ACTION_ALL = "all"

func GetSupportedLanguages() []string {
  return []string { "java", "javascript" }
}

func GetSupportedApps() []string {
  return []string { "vim", "nvim", "code" }
}
