package command

import (
	"log"
	"strings"
	"testing"

	"github.com/Eldius/game-manager-go/config"
)

func TestGetExecutionEnvVars(t *testing.T) {
	cfg := config.ManagerConfig{
		Workspace: "/tmp/test_workspace_env_vars",
		Verbose:   false,
	}
	env := GetExecutionEnvVars(cfg)

	var pathVars []string
	var homeVars []string
	for _, p := range env {
		if strings.HasPrefix(p, "PATH=") {
			pathVars = append(pathVars, p)
		}
		if strings.HasPrefix(p, "HOME=") {
			homeVars = append(homeVars, p)
		}
	}

	log.Println("path:", pathVars)
	if len(pathVars) != 2 {
		t.Errorf("Must have 2 elements starting with 'PATH', but have %d", len(pathVars))
	}

	log.Println("home:", homeVars)
	if len(homeVars) != 2 {
		t.Errorf("Must have 2 elements starting with 'HOME', but have %d", len(homeVars))
	}
}
