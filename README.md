# FZMove

Move like the flash between your projects, folders and files using the terminal. 
Link projects to your favourite IDE/editor and open the project in seconds.

FZMove is a CLI application triggered by `fzmove`. 
If you want a faster command, use an alias to map what you desire.

## CLI 

## This is basic planning. Final README to be ironed out later

* `fzmove` - Root Application


fzmove init - Initalize Dependencies
fzmove - Scans all projects and opens in configured application
fzmove all - Scans all files and folders in the system
    --d - Directory only
        -- t - Change directory with tmux
        -- open - Open folder or file in File Explorer
    --f - Files only ( Opens in default application )
    --copy - Copy path without any action
    --action - When used with `--copy`, copy and action will be done. without copy has no meaning


====


### Configuration

```json
{
  "project_dirs": [ "dir1_path", "dir2_path" ],
  "project_lang": {
    "java": "intellij",
    "go": "nvim",
    "javascript": "vscode",
    "ruby": "custom1"
  },
  "project_custom_apps": {
    "custom1": "lvim"
  }
}
```
 

---

### Planned Feature List

[x] Able to initialize all required dependencies
[] Health check for all required dependencies
[] Browse all directories / files with the option of browsing only directories or files
[] Browse Projects with `.git` as Project Identifier
[] Detect projects based on language quirks := pom.xml means java project, go.mod for golang project and so on [ If not possible, popup when user opens a project for the first time to ask for language]
[] Configure each language with an application to open the project.
[] Folder structure preview
[] File preview with syntax highlighting
[] Path bookmarks
[] CD on steroids
[] Open file on steroids
[] Ability to open projects / directories with custom apps [ vim, emacs etc.. ]
