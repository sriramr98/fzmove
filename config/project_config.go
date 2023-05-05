package config

import (
  "errors"

  "github.com/spf13/viper"
  "gitub.com/sriramr98/fzmove/utils"
)

type ProjectConfig struct {
  Project_dirs []string
  Project_lang map[string]string
  Project_custom_apps map[string]string
}

var project_config *ProjectConfig = nil

func GetAndValidateConfig() (ProjectConfig, error) {

  if project_config != nil {
    return *project_config, nil
  }

  project_dirs := viper.GetStringSlice("project_dirs")
  project_lang := viper.GetStringMapString("project_lang")
  project_custom_apps := viper.GetStringMapString("project_custom_apps")

  if len(project_dirs) == 0 {
    return ProjectConfig{}, errors.New("No project dirs configured")
  }

  supported_langs := utils.GetSupportedLanguages()
  supported_apps := utils.GetSupportedApps()

  for lang, app := range project_lang {

    if !utils.SliceContains(supported_langs, lang) {
      return ProjectConfig{}, errors.New("Invalid Language " + lang)
    }

    if _, ok := project_custom_apps[app]; !utils.SliceContains(supported_apps, app) && !ok {
      return ProjectConfig{}, errors.New("Invalid App " + app)
    }
  }

  return ProjectConfig {
    Project_dirs: project_dirs, 
    Project_lang: project_lang,
    Project_custom_apps: project_custom_apps,
  }, nil

}
