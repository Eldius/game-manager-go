package command

import (
	"os/exec"
	"log"
	"os"
	"fmt"
	"github.com/Eldius/game-manager-go/config"
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
	cmd := &exec.Cmd{
		Path:   scriptPath,
		Args:   execArgs,
		Env:    GetExecutionEnvVars(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	executeCmd(cmd)

}

/*
ExecutePyenvCommand just executes a pyenv command
*/
func ExecutePyenvCommand(args []string) {

	pyenv := filepath.Join(config.GetPyenvBinFolder(), "pyenv")

	execArgs := append([]string{pyenv}, args...)
	cmd := &exec.Cmd{
		Path: pyenv,
		Args: execArgs,
		Env:  GetExecutionEnvVars(),
		//Stdin: os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	executeCmd(cmd)
}

/*
ExecuteShellCommand executes a command
*/
func ExecuteShellCommand(command []string) {
	executeCmd(&exec.Cmd{
		Args: command,
		Env: GetExecutionEnvVars(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
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
