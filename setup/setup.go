package setup

import (
	"os"
	"path/filepath"

	"github.com/Eldius/game-manager-go/config"
)

/*
Setup sets up the environment
*/
func Setup() {
	if !ValidateWorkspaceFolder() {
		SetWorkspaceFolder()
	}

	if !ValidatePyenv() {
		SetPyenv()
	}
	GenerateScripts()
	SetPython()
}

/*
SetWorkspaceFolder Creates the bin folder
*/
func SetWorkspaceFolder() {
	os.MkdirAll(config.WorkspaceFolder(), os.ModePerm)
}

/*
ValidateSetup validates dependency
setup
*/
func ValidateSetup() bool {
	return ValidateWorkspaceFolder() && ValidatePyenv()
}

/*
ValidateWorkspaceFolder validates workspace
folder exists
*/
func ValidateWorkspaceFolder() bool {
	if _, err := os.Stat(config.WorkspaceFolder()); err != nil {
		return false
	}
	return true
}

/*
ValidatePyenv validates dependency
setup
*/
func ValidatePyenv() bool {
	if _, err := os.Stat(filepath.Join(config.WorkspaceFolder(), "pyenv")); err != nil {
		return false
	}
	return true
}
