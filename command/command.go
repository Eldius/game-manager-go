package command

import (
	"fmt"
	"github.com/Eldius/game-manager-go/config"
	"github.com/Eldius/game-manager-go/logger"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

/*
GetExecutionEnvVars generates the env vars to execute
commands or scripts
*/
func GetExecutionEnvVars() []string {
	sysPath, _ := os.LookupEnv("PATH")
	newPath := fmt.Sprintf("PATH=%s:%s", config.GetPyenvBinFolder(), sysPath)
	workspace := config.WorkspaceFolder()
	newUserHome := fmt.Sprintf("HOME=%s", workspace)
	pyenvRoot := fmt.Sprintf("PYENV_ROOT=%s/pyenv", workspace)

	return append(os.Environ(), newPath, newUserHome, pyenvRoot)
}

/*
ExecuteScript executes script by file path
*/
func ExecuteScript(scriptPath string) {
	execArgs := append([]string{scriptPath})

	l := logger.NewLogWriter(logger.DefaultLogger())
	cmd := &exec.Cmd{
		Path:   scriptPath,
		Args:   execArgs,
		Env:    GetExecutionEnvVars(),
		Stdout: l,
		Stderr: l,
	}

	executeCmd(cmd)

}

/*
ExecutePyenvCommand just executes a pyenv command
*/
func ExecutePyenvCommand(args []string) {

	pyenv := filepath.Join(config.GetPyenvBinFolder(), "pyenv")

	execArgs := append([]string{pyenv}, args...)

	l := logger.NewLogWriter(logger.DefaultLogger())
	cmd := &exec.Cmd{
		Path: pyenv,
		Args: execArgs,
		Env:  GetExecutionEnvVars(),
		//Stdin: os.Stdin,
		Stdout: l,
		Stderr: l,
	}

	executeCmd(cmd)
}

/*
ExecuteShellCommand executes a command
*/
func ExecuteShellCommand(command []string) {
	l := logger.NewLogWriter(logger.DefaultLogger())
	executeCmd(&exec.Cmd{
		Args:   command,
		Env:    GetExecutionEnvVars(),
		Stdout: l,
		Stderr: l,
	})
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
