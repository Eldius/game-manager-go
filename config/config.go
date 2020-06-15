package config

import (
	"log"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	pyenvFolder   = "pyenv"
	scriptsFolder = "scripts"
)

/*
ManagerConfig represents the
app configuration
*/
type ManagerConfig struct {
	Verbose   bool
	Workspace string
}

var scripts = map[string]string{
	"install_python_env":             "shell/setup/setup_python_environment.sh",
	"minecraft_ansible_requirements": "ansible/minecraft/roles/requirements.yml",
}

/*
ScriptDef represents a model to
render a script
*/
type ScriptDef struct {
	Name     string
	Template string
	Path     string
}

/*
WorkspaceFolder returns the workspace folder
*/
func (c *ManagerConfig) WorkspaceFolder() string {
	if wsDir, err := homedir.Expand(c.Workspace); err != nil {
		panic(err.Error())
	} else {
		return wsDir
	}
}

/*
GetPyenvFolder returns the pyenv path
*/
func (c *ManagerConfig) GetPyenvFolder() string {
	return filepath.Join(c.Workspace, pyenvFolder)
}

/*
GetPyenvBinFolder returns pyenv bin folder
*/
func (c *ManagerConfig) GetPyenvBinFolder() string {
	return filepath.Join(c.Workspace, pyenvFolder, "bin")
}

/*
GetScriptsFolder returns the scripts folder
*/
func (c *ManagerConfig) GetScriptsFolder() string {
	return filepath.Join(c.Workspace, scriptsFolder)
}

/*
GetAllScripts returns all the script models
*/
func (c *ManagerConfig) GetAllScripts() []ScriptDef {
	var scriptList []ScriptDef
	for k := range scripts {
		scriptList = append(scriptList, c.GetScriptInfo(k))
	}

	return scriptList
}

/*
GetScriptInfo returns
*/
func (c *ManagerConfig) GetScriptInfo(scriptName string) ScriptDef {
	return ScriptDef{
		Name:     scriptName,
		Template: scripts[scriptName],
		Path:     c.GetScriptPath(scriptName),
	}
}

/*
GetScriptPath Returns the path for this script
*/
func (c *ManagerConfig) GetScriptPath(scriptName string) string {
	return filepath.Join(c.GetScriptsFolder(), scripts[scriptName])
}

/*
SaveConfiguration persists the
current configuration
*/
func SaveConfiguration() {
	viper.SafeWriteConfig()
}

/*
GetAppConfig returns the app config
*/
func GetAppConfig() *ManagerConfig {
	var cfg ManagerConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panic(err.Error())
	}

	return &cfg
}
