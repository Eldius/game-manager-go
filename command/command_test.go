package command

import (
	"testing"
	"strings"
	"log"
)

func TestGetExecutionEnvVars(t *testing.T) {
	env := GetExecutionEnvVars()

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
