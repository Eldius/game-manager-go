package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Eldius/game-manager-go/config"
	"github.com/Eldius/game-manager-go/logger"
	"github.com/Eldius/game-manager-go/scripts"
)

/*
GetScriptExecutionEnvVars generates the env vars to execute
scripts
*/
func GetScriptExecutionEnvVars(s scripts.ScriptDef) []string {
	cfg := s.ScriptConfig()
	sysPath, _ := os.LookupEnv("PATH")
	newPath := fmt.Sprintf("PATH=%s:%s", cfg.GetPyenvBinFolder(), sysPath)
	workspace := cfg.Workspace
	newUserHome := fmt.Sprintf("HOME=%s", workspace)
	pyenvRoot := fmt.Sprintf("PYENV_ROOT=%s/pyenv", workspace)

	return append(os.Environ(), newPath, newUserHome, pyenvRoot)
}

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
ExecuteScript executes script by file path
*/
func ExecuteScript(s scripts.ScriptDef) {

	execArgs := append([]string{s.Path})

	l := logger.NewLogWriter(logger.DefaultLogger())
	cmd := &exec.Cmd{
		Path:   s.Path,
		Args:   execArgs,
		Env:    GetScriptExecutionEnvVars(s),
		Stdout: l,
		Stderr: l,
	}

	executeCmd(cmd)

}

/*
ExecuteProvisiningScript executes script by file path
*/
func ExecuteProvisiningScript(s scripts.ScriptDef) {
	execArgs := append([]string{s.Path})

	l := logger.NewLogWriter(logger.DefaultLogger())
	cmd := &exec.Cmd{
		Path:   s.Path,
		Args:   execArgs,
		Env:    GetScriptExecutionEnvVars(s),
		Stdout: l,
		Stderr: l,
	}

	executeCmd(cmd)

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

/*
ExecuteShellCommand executes a command
*/
func ExecuteShellCommand(command []string, cfg config.ManagerConfig) {
	l := logger.NewLogWriter(logger.DefaultLogger())
	executeCmd(&exec.Cmd{
		Args:   command,
		Env:    GetCommandExecutionEnvVars(cfg),
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
