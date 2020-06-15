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

	if cfg.Verbose {
		log.Println(cfg)
	}
	return &cfg
}
