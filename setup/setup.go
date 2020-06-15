package setup

import (
	"os"
	"path/filepath"

	"github.com/Eldius/game-manager-go/config"
)

/*
Setup sets up the environment
*/
func Setup(cfg config.ManagerConfig) {
	GenerateScripts(cfg)
	if !ValidateWorkspaceFolder(cfg) {
		SetWorkspaceFolder(cfg)
	}

	if !ValidatePyenv(cfg) {
		SetPyenv(cfg)
	}
	SetPythonEnv(cfg)
	config.SaveConfiguration()
}

/*
SetWorkspaceFolder Creates the bin folder
*/
func SetWorkspaceFolder(cfg config.ManagerConfig) {
	os.MkdirAll(cfg.Workspace, os.ModePerm)
}

/*
ValidateSetup validates dependency
setup
*/
func ValidateSetup(cfg config.ManagerConfig) bool {
	return ValidateWorkspaceFolder(cfg) && ValidatePyenv(cfg)
}

/*
ValidateWorkspaceFolder validates workspace
folder exists
*/
func ValidateWorkspaceFolder(cfg config.ManagerConfig) bool {
	if _, err := os.Stat(cfg.Workspace); err != nil {
		return false
	}
	return true
}

/*
ValidatePyenv validates dependency
setup
*/
func ValidatePyenv(cfg config.ManagerConfig) bool {
	if _, err := os.Stat(filepath.Join(cfg.Workspace, "pyenv")); err != nil {
		return false
	}
	return true
}
