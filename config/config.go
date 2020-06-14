package config

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	pyenvFolder   = "pyenv"
	scriptsFolder = "scripts"
)

var scripts = map[string]string{
	"install_python_env": "shell/setup/setup_python_environment.sh",
	"ansible_requirements": "ansible/roles/requirements.yml",
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
func WorkspaceFolder() string {
	if wsDir, err := homedir.Expand(viper.GetString("workspace")); err != nil {
		panic(err.Error())
	} else {
		return wsDir
	}
}

/*
GetPyenvFolder returns the pyenv path
*/
func GetPyenvFolder() string {
	return filepath.Join(WorkspaceFolder(), pyenvFolder)
}

/*
GetPyenvBinFolder returns pyenv bin folder
*/
func GetPyenvBinFolder() string {
	return filepath.Join(WorkspaceFolder(), pyenvFolder, "bin")
}

/*
GetScriptsFolder returns the scripts folder
*/
func GetScriptsFolder() string {
	return filepath.Join(WorkspaceFolder(), scriptsFolder)
}

/*
InstallPythonEnvScript returns the install
Python script path
*/
//func InstallPythonEnvScript() string {
//	return filepath.Join(GetScriptsFolder(), scripts["install_python_env"])
//}

/*
GetAllScripts returns all the script models
*/
func GetAllScripts() []ScriptDef {
	var scriptList []ScriptDef
	for k := range scripts {
		scriptList = append(scriptList, GetScriptInfo(k))
	}

	return scriptList
}

/*
GetScriptInfo returns
*/
func GetScriptInfo(scriptName string) ScriptDef {
	return ScriptDef{
		Name:     scriptName,
		Template: scripts[scriptName],
		Path:     GetScriptPath(scriptName),
	}
}

/*
GetScriptPath Returns the path for this script
*/
func GetScriptPath(scriptName string) string {
	return filepath.Join(GetScriptsFolder(), scripts[scriptName])
}

/*
SaveConfiguration persists the
current configuration
*/
func SaveConfiguration() {
	viper.SafeWriteConfig()
}
