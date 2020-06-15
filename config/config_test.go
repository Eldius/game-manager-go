package config

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestConfiLoad(t *testing.T) {
	cfgFile, err := filepath.Abs("./samples/config.yml")
	if err != nil {
		t.Error(err)
	}
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}

	t.Log(fmt.Sprintf("config file at: %s", cfgFile))
	cfg := GetAppConfig()

	if cfg.Workspace != "/tmp/test_cfg" {
		t.Errorf("Workspace should be '/tmp/test_cfg', but was '%s'", cfg.Workspace)
	}
	if !cfg.Verbose {
		t.Errorf("Verbose should be 'true', but was '%v'", cfg.Verbose)
	}
}
