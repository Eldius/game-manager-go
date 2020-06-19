package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Eldius/game-manager-go/config"
	"github.com/Eldius/game-manager-go/logger"
)

/*
GetCommandExecutionEnvVars generates the env vars to execute
commands
*/
func GetCommandExecutionEnvVars(cfg config.ManagerConfig) []string {
	sysPath, _ := os.LookupEnv("PATH")
	newPath := fmt.Sprintf("PATH=%s:%s", cfg.GetPyenvBinFolder(), sysPath)
	workspace := cfg.Workspace
	newUserHome := fmt.Sprintf("HOME=%s", workspace)
	pyenvRoot := fmt.Sprintf("PYENV_ROOT=%s/pyenv", workspace)

	return append(os.Environ(), newPath, newUserHome, pyenvRoot)
}

/*
ExecutePyenvCommand just executes a pyenv command
*/
func ExecutePyenvCommand(args []string, cfg config.ManagerConfig) {

	pyenv := filepath.Join(cfg.GetPyenvBinFolder(), "pyenv")

	execArgs := append([]string{pyenv}, args...)

	l := logger.NewLogWriter(logger.DefaultLogger())
	cmd := &exec.Cmd{
		Path: pyenv,
		Args: execArgs,
		Env:  GetCommandExecutionEnvVars(cfg),
		//Stdin: os.Stdin,
		Stdout: l,
		Stderr: l,
	}

	executeCmd(cmd)
}

func executeCmd(cmd *exec.Cmd) {
	log.Println("cmd:", cmd.String())

	log.Println()
	log.Println("**********")
	if err := cmd.Run(); err != nil {
		log.Println("---")
		log.Println("Failed to install python")
		log.Println(err.Error())
	}
	log.Println("**********")
}
