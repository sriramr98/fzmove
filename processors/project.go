package processors

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"

	"gitub.com/sriramr98/fzmove/config"
	"gitub.com/sriramr98/fzmove/core"
)

type ProjectProcessor struct {}

func NewProjectProcessor() ProjectProcessor {
  return ProjectProcessor{}
}

func (p ProjectProcessor) process() error {
  project_config, err := config.GetAndValidateConfig()
  if err != nil {
    return err
  }

  // Scan each folder  projects
  project_dirs := extractProjectDirs(project_config.Project_dirs)
  // Feed list to FzF
  
  project_path, err := core.ProcessFzF(strings.Join(project_dirs, "\n"))
  if err != nil {
    return err
  }

  // Depending on selection and identified project type, open project
  project_type := identifyProjectType(project_path)

  configured_app := extractAppForType(project_type, project_config)

  core.Move(project_path, configured_app)

  return nil
}

func extractProjectDirs(dirs []string) []string {
  all_project_dirs := []string{}

  for _, root_dir := range dirs {
    result, err := os.ReadDir(root_dir)
    if err != nil {
      fmt.Printf("Unable to read directory %s, error: %s \n", root_dir, err)
      continue
    }

    for _, file := range result {
      if file.IsDir() {
        if is_valid_project := validateProject(file); is_valid_project {
          all_project_dirs = append(all_project_dirs, path.Join(root_dir, file.Name()))
        }
      }
    }
  }

  return all_project_dirs
}

func validateProject(file fs.DirEntry) bool {
  // TODO: Implement
  return true
}

func identifyProjectType(path string) string {
  return ""
}

func extractAppForType(projectType string, projectConfig config.ProjectConfig) string {
  return ""
}
