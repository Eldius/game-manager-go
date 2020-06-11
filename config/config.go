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
func InstallPythonEnvScript() string {
	return filepath.Join(GetScriptsFolder(), "setup_python_environment.sh")
}
