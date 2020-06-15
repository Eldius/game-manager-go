package setup

import (
	"log"
	"os"
	"runtime"

	"github.com/Eldius/game-manager-go/command"
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
func SetPyenv(cfg config.ManagerConfig) {
	if runtime.GOOS == "linux" {
		log.Println("Cloning pyenv...")
		if repo, err := clone(pyenvRepo, cfg.GetPyenvFolder()); err != nil {
			log.Panic(err.Error())
		} else {
			repo.Fetch(&git.FetchOptions{
				RemoteName: git.DefaultRemoteName,
				Progress:   os.Stdout,
			})
		}
	} else if runtime.GOOS == "windows" {
		log.Println("[not implemented yet] Cloning pyenv-win...")
		os.Exit(1)
	}
}

/*
SetPythonEnv sets up the Python environment
*/
func SetPythonEnv(cfg config.ManagerConfig) {
	log.Println("seting up ansible")

	command.ExecuteScript(cfg.GetScriptPath("install_python_env"), cfg)
}
