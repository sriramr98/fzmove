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
    
