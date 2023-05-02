package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
)


func InitConfig() {
  viper.SetEnvPrefix("FZMV")


  home_dir, err := os.UserHomeDir()
  if err != nil {
    log.Fatal(err)
  } 

  viper.SetDefault("config_path", path.Join(home_dir, "/.config/fzmove/"))

  if err := viper.BindEnv("config_path"); err != nil {
    fmt.Println(err.Error());
  }

  config_path := viper.GetString("config_path")

  viper.AddConfigPath(config_path)
  viper.SetConfigName("config")
  viper.ReadInConfig()
}

func GetProjectDirs() []string {
  return viper.GetStringSlice("project_dirs")
}
