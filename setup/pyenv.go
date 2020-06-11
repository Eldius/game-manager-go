package setup

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Eldius/game-manager-go/config"
	"github.com/go-git/go-git/v5"
)

const (
	pyenvRepo = "https://github.com/pyenv/pyenv.git"
)

func clone(repo string, dest string) (r *git.Repository, err error) {
	r, err = git.PlainClone(dest, false, &git.CloneOptions{
		URL:      repo,
		Progress: os.Stdout,
	})
	return
}

/*
SetPyenv clones pyenv repository
*/
func SetPyenv() {
	if runtime.GOOS == "linux" {
		log.Println("Cloning pyenv...")
		if repo, err := clone(pyenvRepo, config.GetPyenvFolder()); err != nil {
			log.Panic(err.Error())
		} else {
			repo.Fetch(&git.FetchOptions{
				RemoteName: git.DefaultRemoteName,
				Progress:   os.Stdout,
			})
		}
		SetPython()
		SetAnsible()
	} else if runtime.GOOS == "windows" {
		log.Println("[not implemented yet] Cloning pyenv-win...")
		os.Exit(1)
	}
}

/*
SetPython installs ansible
*/
func SetPython() {
	PyenvExecuteCommand([]string{"help"})
	executeScript(config.GetInstallPythonFile())
}

func executeScript(scriptPath string) {
	execArgs := append([]string{scriptPath})
	cmd := &exec.Cmd{
		Path:   scriptPath,
		Args:   execArgs,
		Env:    getExecutionEnvVars(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	log.Println("cmd:", cmd.String())

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("pyenv path:\n", cmd.Env)
	log.Println()
	log.Println("---")
	if err := cmd.Run(); err != nil {
		log.Println("---")
		log.Panic(err.Error())
	}
	log.Println("---")

}

/*
Test just test
*/
func Test(args []string) {
	PyenvExecuteCommand(args)
}

func getExecutionEnvVars() []string {
	sysPath, _ := os.LookupEnv("PATH")
	newPath := fmt.Sprintf("PATH=%s:%s", config.GetPyenvBinFolder(), sysPath)
	workspace := config.WorkspaceFolder()
	newUserHome := fmt.Sprintf("HOME=%s", workspace)
	pyenvRoot := fmt.Sprintf("PYENV_ROOT=%s/pyenv", workspace)

	return append(os.Environ(), newPath, newUserHome, pyenvRoot)
}

/*
PyenvExecuteCommand just test
*/
func PyenvExecuteCommand(args []string) {

	pyenv := filepath.Join(config.GetPyenvBinFolder(), "pyenv")

	execArgs := append([]string{pyenv}, args...)
	cmd := &exec.Cmd{
		Path: pyenv,
		Args: execArgs,
		Env:  getExecutionEnvVars(),
		//Stdin: os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	log.Println("cmd:", cmd.String())

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println()
	log.Println("---")
	if err := cmd.Run(); err != nil {
		log.Println("---")
		log.Panic(err.Error())
	}
	log.Println("---")
}

/*
ShellTest just test
*/
func ShellTest() {
	cmd := exec.Command("env")
	log.Println("cmd:", cmd.String())

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	sysPath, _ := os.LookupEnv("PATH")

	newPath := fmt.Sprintf("PATH=\"%s:%s\"", config.GetPyenvBinFolder(), sysPath)
	newUserHome := fmt.Sprintf("HOME=%s", config.WorkspaceFolder())
	cmd.Env = []string{newPath, newUserHome}
	log.Println()
	log.Println("---")
	if err := cmd.Run(); err != nil {
		log.Println("---")
		log.Panic(err.Error())
	}
	log.Println("---")
}

/*
SetAnsible installs ansible
*/
func SetAnsible() {
	log.Println("seting up ansible")
	cmd := exec.Command("pip install ansible")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	var env []string
	pyenvFolder := config.GetPyenvFolder()
	pyenvBinFolder := filepath.Join(pyenvFolder, "bin")
	sysPath, _ := os.LookupEnv("PATH")
	newPath := strings.Join([]string{pyenvBinFolder, sysPath}, ":")
	log.Println("path:", newPath)
	env = append(env, fmt.Sprintf("PATH=%s", newPath), fmt.Sprintf("HOME=%s", pyenvFolder))
	cmd.Env = env

	if err := cmd.Run(); err != nil {
		log.Panic(err.Error())
	}
}
